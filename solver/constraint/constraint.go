package constraint

import (
	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/score"
)

type ConstraintType int

const (
	HARD ConstraintType = iota // 硬约束
	SOFT                       // 软约束
)

type Constraint struct {
	// 约束名称
	Name string
	// 约束权重
	Weight int
	// 约束类型
	Type ConstraintType
	// 约束匹配函数
	MatchFunc func(solution api.ISolution) bool
}

func NewConstraint(options ...func(*Constraint)) *Constraint {
	constraint := &Constraint{}
	for _, option := range options {
		option(constraint)
	}
	return constraint
}

func WithName(name string) func(*Constraint) {
	return func(constraint *Constraint) {
		constraint.Name = name
	}
}

func WithWeight(weight int) func(*Constraint) {
	return func(constraint *Constraint) {
		constraint.Weight = weight
	}
}

func WithType(constraintType ConstraintType) func(*Constraint) {
	return func(constraint *Constraint) {
		constraint.Type = constraintType
	}
}

func WithMatchFunc(matchFunc func(solution api.ISolution) bool) func(*Constraint) {
	return func(constraint *Constraint) {
		constraint.MatchFunc = matchFunc
	}
}

func (c *Constraint) GetScore() api.IScore {
	switch c.Type {
	case HARD:
		return score.NewHardSoftScore(c.Weight, 0, 0)
	case SOFT:
		return score.NewHardSoftScore(0, c.Weight, 0)
	default:
		return score.NewHardSoftScore(0, 0, 0)
	}
}

func (c *Constraint) Match(solution api.ISolution) bool {
	return c.MatchFunc(solution)
}

func (c *Constraint) GetWeight() int {
	return c.Weight
}
