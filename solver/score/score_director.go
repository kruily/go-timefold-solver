package score

import "github.com/kruily/go-timefold-solver/solver/api"

type ScoreDirector struct {
	calculator           *ScoreCalulator
	increamentCalculator *IncrementalScoreCalculator
	solution             api.ISolution
	useIncreament        bool
}

func NewScoreDirector(calculator *ScoreCalulator, constraintManager api.IConstraintConfigure) *ScoreDirector {
	return &ScoreDirector{
		calculator:           calculator,
		increamentCalculator: NewIncrementalScoreCalculator(constraintManager),
		useIncreament:        false,
	}
}

func (s *ScoreDirector) Calculate(solution api.ISolution) api.IScore {
	if s.useIncreament {
		// 增量计算
		s.increamentCalculator.solution = solution
		return s.increamentCalculator.scoreCache
	}
	// 完全计算
	return s.calculator.Calculate(solution)
}

func (s *ScoreDirector) BeforeVariableChanged(planningVariable api.IPlanningVariable) {
	if s.useIncreament {
		s.increamentCalculator.BeforeVariableChange(planningVariable)
	}
}

func (s *ScoreDirector) AfterVariableChanged(planningVariable api.IPlanningVariable) {
	if s.useIncreament {
		s.increamentCalculator.AfterVariableChange(planningVariable)
	}
}

func (s *ScoreDirector) GetWorkingSolution() api.ISolution {
	return s.solution
}

func (s *ScoreDirector) SetWorkingSolution(solution api.ISolution) {
	s.solution = solution
}

func (s *ScoreDirector) SetUseIncreament(useIncreament bool) {
	s.useIncreament = useIncreament
}
