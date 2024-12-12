package score

import (
	"github.com/kruily/go-timefold-solver/solver/api"
)

type ScoreCalulator struct {
	constraintManager api.IConstraintConfigure
	constraints       []api.IConstraint
}

func NewScoreCalculator(constraintManager api.IConstraintConfigure) *ScoreCalulator {
	return &ScoreCalulator{constraintManager: constraintManager}
}

func (s *ScoreCalulator) Calculate(solution api.ISolution) api.IScore {
	hardScore := 0
	softScore := 0
	for _, constraint := range s.constraintManager.GetConstraints() {
		if constraint.Match(solution) {
			score := constraint.GetScore().(*HardSoftScore)
			hardScore += score.hardScore * constraint.GetWeight()
			softScore += score.softScore * constraint.GetWeight()
		}
	}
	return NewHardSoftScore(hardScore, softScore, 0)
}
