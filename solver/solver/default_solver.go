package solver

import (
	"context"
	"math"
	"math/rand/v2"
	"sort"
	"sync"
	"time"

	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/config"
	"github.com/kruily/go-timefold-solver/solver/move"
	"github.com/kruily/go-timefold-solver/solver/tabu"
)

type DefaultSolver struct {
	config        *config.SolverConfig
	scoreDirector api.IScoreDirector
	moveSelector  move.MoveSelector
	tabuAcceptor  *tabu.TabuSearchAcceptor
	currentMove   api.IMove

	terminated   bool
	terminateMu  sync.Mutex
	bestSolution api.ISolution
	bestScore    api.IScore

	ctx    context.Context
	cancel context.CancelFunc
}

func NewDefaultSolver(cfg *config.SolverConfig, scoreDirector api.IScoreDirector) *DefaultSolver {
	ctx, cancel := context.WithCancel(context.Background())
	solver := &DefaultSolver{
		config:        cfg,
		scoreDirector: scoreDirector,
		ctx:           ctx,
		cancel:        cancel,
	}
	// 创建禁忌搜索接受器
	aspirationConfig := config.NewAspirationConfig(
		cfg.LocalSearchConfig.TabuSearchConfig.AspirationCriteria,
		cfg.LocalSearchConfig.TabuSearchConfig.TimeLimit,
		cfg.LocalSearchConfig.TabuSearchConfig.MaxFrequency,
	)
	solver.tabuAcceptor = tabu.NewTabuSearchAcceptor(
		cfg.LocalSearchConfig.TabuSearchConfig.MaxFrequency,
		cfg.LocalSearchConfig.TabuSearchConfig.MaxFrequency,
		aspirationConfig,
	)
	solver.moveSelector = move.NewDefaultMoveSelector(cfg, scoreDirector)
	return solver
}

func (s *DefaultSolver) Solve(problem api.ISolution) (api.ISolution, error) {
	s.init(problem)
	// 设置工作解
	s.scoreDirector.SetWorkingSolution(problem)

	if s.config.TimeLimit > 0 {
		ctx, cancel := context.WithTimeout(s.ctx, time.Duration(s.config.TimeLimit)*time.Second)
		defer cancel()
		s.ctx = ctx
	}

	// 构造初始解
	solution := s.constructInitialSolution(problem)
	s.updateBestSolution(solution)

	// 使用局部搜索进行改进解
	if s.config.LocalSearch {
		solution = s.localSearch(solution)
	}
	return s.bestSolution, nil
}

func (s *DefaultSolver) Stop() {
	s.terminateMu.Lock()
	defer s.terminateMu.Unlock()
	s.terminated = true
	s.cancel()
}

func (s *DefaultSolver) IsTerminated() bool {
	s.terminateMu.Lock()
	defer s.terminateMu.Unlock()
	return s.terminated
}

func (s *DefaultSolver) GetBestSolution() api.ISolution {
	return s.bestSolution
}

func (s *DefaultSolver) init(problem api.ISolution) {
	s.terminateMu.Lock()
	s.terminated = false
	s.terminateMu.Unlock()

	if s.tabuAcceptor != nil {
		s.tabuAcceptor.Clear()
	}
	initailScore := s.scoreDirector.Calculate(problem)
	problem.SetScore(initailScore)
	s.bestSolution = problem
	s.bestScore = initailScore

	s.currentMove = nil
}

func (s *DefaultSolver) updateBestSolution(solution api.ISolution) {
	score := s.scoreDirector.Calculate(solution)
	solution.SetScore(score)
	if s.bestScore == nil || score.CompareTo(s.bestScore) > 0 {
		s.bestSolution = solution
		s.bestScore = score
	}
}

func (s *DefaultSolver) localSearch(solution api.ISolution) api.ISolution {
	currentSolution := solution
	currentScore := s.scoreDirector.Calculate(currentSolution)

	lsConfig := s.config.LocalSearchConfig
	temperature := lsConfig.InitialTemperature

	stepCount := 0
	lastImprovement := 0
	for !s.IsTerminated() {
		if s.checkTermination(stepCount, lastImprovement) {
			break
		}
		move := s.selectMove(currentSolution)
		if move == nil {
			break
		}
		s.currentMove = move

		s.scoreDirector.BeforeVariableChanged(nil) // 需要正确实现
		move.Execute(currentSolution)
		s.scoreDirector.AfterVariableChanged(nil) // 需要正确实现
		newScore := s.scoreDirector.Calculate(currentSolution)
		accept := s.acceptMove(currentScore, newScore, temperature, &lsConfig)
		if accept {
			currentScore = newScore
			s.updateBestSolution(currentSolution)
			if newScore.CompareTo(s.bestScore) > 0 {
				lastImprovement = 0
			}
		} else {
			move.Undo(currentSolution)
		}
		stepCount++
		temperature *= lsConfig.CoolingRate
	}
	return s.bestSolution
}

func (s *DefaultSolver) checkTermination(stepCount int, lastImprovement int) bool {
	termConfig := s.config.Termination
	if termConfig.StepCountLimit > 0 && stepCount >= termConfig.StepCountLimit {
		return true
	}
	if termConfig.UnimprovedStepCountLimit > 0 && (stepCount-lastImprovement) >= termConfig.UnimprovedStepCountLimit {
		return true
	}
	if termConfig.BestScoreLimit != nil && s.bestScore.CompareTo(termConfig.BestScoreLimit) >= 0 {
		return true
	}
	return false
}

func (s *DefaultSolver) selectMove(solution api.ISolution) api.IMove {
	return s.moveSelector.SelectMove(solution)
}

func (s *DefaultSolver) acceptMove(currentScore, newScore api.IScore, temperature float64, lsConfig *config.LocalSearchConfig) bool {
	switch lsConfig.Type {
	case config.LocalSearchTypeSimulatedAnnealing:
		return s.acceptSimulatedAnnealing(currentScore, newScore, temperature)
	case config.LocalSearchTypeTabuSearch:
		return s.acceptTabuSearch(newScore, lsConfig)
	default:
		return newScore.CompareTo(currentScore) > 0
	}
}

func (s *DefaultSolver) acceptSimulatedAnnealing(currentScore, newScore api.IScore, temperature float64) bool {
	if newScore.CompareTo(currentScore) >= 0 {
		return true
	}
	delta := float64(newScore.CompareTo(currentScore))
	probability := math.Exp(delta / temperature)
	return rand.Float64() < probability
}

func (s *DefaultSolver) acceptTabuSearch(newScore api.IScore, lsConfig *config.LocalSearchConfig) bool {
	move := s.currentMove
	accept, err := s.tabuAcceptor.Accept(move, newScore)
	if err != nil {
		return false
	}
	if accept {
		err = s.tabuAcceptor.RecordMove(move, newScore)
		if err != nil {
			return false
		}
	}
	return accept
}

func (s *DefaultSolver) constructInitialSolution(problem api.ISolution) api.ISolution {
	switch s.config.ConstructionHeuristic {
	case "FIRST_FIT":
		return s.firstFit(problem)
	case "FIRST_FIT_DECREASING":
		return s.firstFitDecreasing(problem)
	default:
		return s.firstFit(problem)
	}
}

// firstFit 实现最先适应构造法
func (s *DefaultSolver) firstFit(problem api.ISolution) api.ISolution {
	// 获取所有规划实体
	entities := s.getPlanningEntities(problem)

	// 对每个实体进行赋值
	for _, entity := range entities {
		// 获取实体的所有规划变量
		variables := entity.GetPlanningVariables()

		for _, variable := range variables {
			// 获取变量的值域
			valueRange := variable.GetValueRange()
			iterator := valueRange.CreateIterator()

			// 找到第一个可行值
			for iterator.HasNext() {
				value := iterator.Next()
				variable.SetValue(value)

				// 计算当前解的得分
				score := s.scoreDirector.Calculate(problem)
				if score.IsFeasible() {
					break
				}
			}
		}
	}

	return problem
}

// firstFitDecreasing 实现最先适应递减构造法
func (s *DefaultSolver) firstFitDecreasing(problem api.ISolution) api.ISolution {
	// 获取所有规划实体
	entities := s.getPlanningEntities(problem)

	// 对实体进行排序（按照某种启发式规则）
	sort.Slice(entities, func(i, j int) bool {
		return s.compareEntities(entities[i], entities[j])
	})

	// 使用排序后的实体列表执行firstFit
	for _, entity := range entities {
		variables := entity.GetPlanningVariables()

		for _, variable := range variables {
			valueRange := variable.GetValueRange()
			iterator := valueRange.CreateIterator()

			bestValue := iterator.Next()
			bestScore := s.evaluateAssignment(problem, variable, bestValue)

			// 尝试所有可能的值，选择最好的
			for iterator.HasNext() {
				value := iterator.Next()
				score := s.evaluateAssignment(problem, variable, value)

				if score.CompareTo(bestScore) > 0 {
					bestValue = value
					bestScore = score
				}
			}

			variable.SetValue(bestValue)
		}
	}

	return problem
}

// getPlanningEntities 获取问题中的所有规划实体
func (s *DefaultSolver) getPlanningEntities(solution api.ISolution) []api.IPlanningEntity {
	// 从解决方案中提取所有规划实体
	var entities []api.IPlanningEntity

	// 获取问题事实
	facts := solution.GetProblemFacts()
	for _, fact := range facts {
		if entity, ok := fact.(api.IPlanningEntity); ok {
			entities = append(entities, entity)
		}
	}

	return entities
}

// compareEntities 比较两个实体的优先级
func (s *DefaultSolver) compareEntities(a, b api.IPlanningEntity) bool {
	// 获取实体的规划变量数量作为优先级依据
	aVars := len(a.GetPlanningVariables())
	bVars := len(b.GetPlanningVariables())

	// 变量越多，优先级越高
	if aVars != bVars {
		return aVars > bVars
	}

	// TODO: 可以添加更多的启发式规则
	return false
}

// evaluateAssignment 评估单个赋值的效果
func (s *DefaultSolver) evaluateAssignment(solution api.ISolution, variable api.IPlanningVariable, value interface{}) api.IScore {
	// 保存原始值
	originalValue := variable.GetValue()

	// 尝试新值
	s.scoreDirector.BeforeVariableChanged(variable)
	variable.SetValue(value)
	s.scoreDirector.AfterVariableChanged(variable)

	// 计算得分
	score := s.scoreDirector.Calculate(solution)

	// 恢复原始值
	s.scoreDirector.BeforeVariableChanged(variable)
	variable.SetValue(originalValue)
	s.scoreDirector.AfterVariableChanged(variable)

	return score
}
