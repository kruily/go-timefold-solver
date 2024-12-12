package constraint

import "github.com/kruily/go-timefold-solver/solver/api"

type ConstraintManager struct {
	constraints []*Constraint
}

func NewConstraintManager() *ConstraintManager {
	return &ConstraintManager{constraints: make([]*Constraint, 0)}
}

func (c *ConstraintManager) AddConstraint(constraint *Constraint) {
	c.constraints = append(c.constraints, constraint)
}

func (c *ConstraintManager) GetConstraints() []api.IConstraint {
	result := make([]api.IConstraint, 0)
	for i, c := range c.constraints {
		result[i] = c
	}
	return result
}
