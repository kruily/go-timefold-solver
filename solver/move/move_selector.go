package move

import (
	"math/rand"
	"time"

	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/config"
)

type MoveSelector interface {
	SelectMove(solution api.ISolution) api.IMove
	Reset()
}

type DefaultMoveSelector struct {
	config        *config.SolverConfig
	scoreDirector api.IScoreDirector
	random        *rand.Rand
}

func NewDefaultMoveSelector(config *config.SolverConfig, scoreDirector api.IScoreDirector) *DefaultMoveSelector {
	return &DefaultMoveSelector{
		config:        config,
		scoreDirector: scoreDirector,
		random:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *DefaultMoveSelector) SelectMove(solution api.ISolution) api.IMove {
	switch s.config.MoveSelector {
	case config.MOVE_SELECTOR_FIRST_FIT:
		return s.selectFirstFitMove(solution)
	case config.MOVE_SELECTOR_BEST_FIT:
		return s.selectBestFitMove(solution)
	case config.MOVE_SELECTOR_CHANGE:
		return s.selectChangeMove(solution)
	case config.MOVE_SELECTOR_CHAINED:
		return s.selectChainedMove(solution)
	default:
		return s.selectFirstFitMove(solution)
	}
}

func (s *DefaultMoveSelector) selectFirstFitMove(solution api.ISolution) api.IMove {
	entities := s.getPlanningEntities(solution)
	if len(entities) < 2 {
		return nil
	}
	for i := 0; i < len(entities); i++ {
		for j := i + 1; j < len(entities); j++ {

			vars1 := entities[i].GetPlanningVariables()
			vars2 := entities[j].GetPlanningVariables()
			for _, v1 := range vars1 {
				move := NewSwapMove(entities[i], entities[j], v1, vars2[0], s.scoreDirector)
				if s.isFeasibleMove(solution, move) {
					return move
				}
			}
		}
	}
	return nil
}

func (s *DefaultMoveSelector) selectBestFitMove(solution api.ISolution) api.IMove {
	entities := s.getPlanningEntities(solution)
	if len(entities) < 2 {
		return nil
	}
	var bestMove api.IMove
	var bestScore api.IScore
	for i := 0; i < len(entities); i++ {
		for j := i + 1; j < len(entities); j++ {
			vars1 := entities[i].GetPlanningVariables()
			vars2 := entities[j].GetPlanningVariables()
			for _, v1 := range vars1 {
				for _, v2 := range vars2 {
					move := NewSwapMove(entities[i], entities[j], v1, v2, s.scoreDirector)
					if !s.isFeasibleMove(solution, move) {
						continue
					}
					score := s.evaluateMove(move, solution)
					if bestMove == nil || score.CompareTo(bestScore) > 0 {
						bestMove = move
						bestScore = score
					}
				}
			}
		}
	}
	return bestMove
}

func (s *DefaultMoveSelector) selectRandomMove(solution api.ISolution) api.IMove {
	entities := s.getPlanningEntities(solution)
	if len(entities) < 2 {
		return nil
	}
	maxAttempts := 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		i := s.random.Intn(len(entities))
		j := s.random.Intn(len(entities))
		if i == j {
			continue
		}
		vars1 := entities[i].GetPlanningVariables()
		vars2 := entities[j].GetPlanningVariables()
		if len(vars1) == 0 || len(vars2) == 0 {
			continue
		}
		v1 := vars1[s.random.Intn(len(vars1))]
		v2 := vars2[s.random.Intn(len(vars2))]
		move := NewSwapMove(entities[i], entities[j], v1, v2, s.scoreDirector)
		if s.isFeasibleMove(solution, move) {
			return move
		}
	}
	return nil
}

func (s *DefaultMoveSelector) selectChainedMove(solution api.ISolution) api.IMove {
	moves := make([]api.IMove, 0)

	move1 := s.selectRandomMove(solution)
	if move1 == nil {
		return nil
	}
	moves = append(moves, move1)
	if s.random.Float64() < 0.5 {
		move2 := s.selectRandomMove(solution)
		if move2 != nil {
			moves = append(moves, move2)
		}
	}
	return NewChainMove(moves, s.scoreDirector)
}

func (s *DefaultMoveSelector) selectChangeMove(solution api.ISolution) api.IMove {
	entities := s.getPlanningEntities(solution)

	for _, entity := range entities {
		variables := entity.GetPlanningVariables()
		for _, variable := range variables {
			valueRange := variable.GetValueRange()
			iterator := valueRange.CreateIterator()
			for iterator.HasNext() {
				value := iterator.Next()
				if value == variable.GetValue() {
					continue
				}
				mov := NewChangeMove(entity, variable, value, s.scoreDirector)
				if s.isFeasibleMove(solution, mov) {
					return mov
				}
			}
		}
	}
	return nil
}

func (s *DefaultMoveSelector) isFeasibleMove(solution api.ISolution, move api.IMove) bool {
	move.Execute(solution)
	score := s.scoreDirector.Calculate(solution)
	feasible := score.IsFeasible()
	move.Undo(solution)
	return feasible
}

func (s *DefaultMoveSelector) evaluateMove(move api.IMove, solution api.ISolution) api.IScore {
	move.Execute(solution)
	score := s.scoreDirector.Calculate(solution)
	move.Undo(solution)
	return score
}

func (s *DefaultMoveSelector) getPlanningEntities(solution api.ISolution) []api.IPlanningEntity {
	var entities []api.IPlanningEntity
	facts := solution.GetProblemFacts()
	for _, fact := range facts {
		if entity, ok := fact.(api.IPlanningEntity); ok {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (s *DefaultMoveSelector) Reset() {
	s.random = rand.New(rand.NewSource(s.config.RandomSeed))
}
