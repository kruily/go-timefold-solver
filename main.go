package main

import (
	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/config"
	"github.com/kruily/go-timefold-solver/solver/constraint"
	"github.com/kruily/go-timefold-solver/solver/score"
	"github.com/kruily/go-timefold-solver/solver/solver"
)

func main() {
	config := config.NewDefalutSolverConfig()
	constraintManager := constraint.NewConstraintManager()
	constraintManager.AddConstraint(constraint.NewConstraint(
		constraint.WithName("constraint1"),
		constraint.WithType(constraint.HARD),
		constraint.WithMatchFunc(func(solution api.ISolution) bool {
			return true
		}),
	))
	calculator := score.NewScoreCalculator(constraintManager)
	scoreDirector := score.NewScoreDirector(calculator, nil)
	solver := solver.NewDefaultSolver(config, scoreDirector)
	solver.Solve(nil)
}
