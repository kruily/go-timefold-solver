package score

import (
	"sync"

	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/solution"
)

type IncrementalScoreCalculator struct {
	solution          api.ISolution
	scoreCache        api.IScore
	constraintManager api.IConstraintConfigure
	mu                sync.Mutex

	dirtyEntities map[api.IPlanningEntity]struct{}
	dirtyVars     map[api.IPlanningVariable]struct{}
}

func NewIncrementalScoreCalculator(constraintManager api.IConstraintConfigure) *IncrementalScoreCalculator {
	return &IncrementalScoreCalculator{
		constraintManager: constraintManager,
		dirtyEntities:     make(map[api.IPlanningEntity]struct{}),
		dirtyVars:         make(map[api.IPlanningVariable]struct{}),
	}
}

func (c *IncrementalScoreCalculator) BeforeVariableChange(variable api.IPlanningVariable) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 将变量标记为脏
	c.dirtyVars[variable] = struct{}{}
	if entity := c.findEntityForVariable(variable); entity != nil {
		c.dirtyEntities[entity] = struct{}{}
	}
}

func (c *IncrementalScoreCalculator) AfterVariableChange(variable api.IPlanningVariable) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 重新计算受影响的约束
	c.recalculateAffectedConstraints()
}

func (c *IncrementalScoreCalculator) findEntityForVariable(variable api.IPlanningVariable) api.IPlanningEntity {
	facts := c.solution.GetProblemFacts()

	for _, fact := range facts {
		if entity, ok := fact.(api.IPlanningEntity); ok {
			for _, v := range entity.GetPlanningVariables() {
				if v == variable {
					return entity
				}
			}
		}
	}

	return nil
}

func (c *IncrementalScoreCalculator) recalculateAffectedConstraints() {
	// 只重新计算受影响的约束
	affectedConstraints := c.calulateScoreForDirtyEntities()
	if c.scoreCache == nil {
		c.scoreCache = affectedConstraints
	} else {
		// TODO：实现分数的增量更新
		c.scoreCache = affectedConstraints
	}
	c.clearDirtyFlags()
}

func (c *IncrementalScoreCalculator) calulateScoreForDirtyEntities() api.IScore {
	sub := c.createSubSolution()
	var totalScore api.IScore

	for _, constraint := range c.constraintManager.GetConstraints() {
		if constraint.Match(sub) {
			score := constraint.GetScore()
			if totalScore == nil {
				totalScore = score
			} else {
				// 假设分数实现了 Add 方法
				if hardSoftScore, ok := score.(*HardSoftScore); ok {
					totalScore = hardSoftScore.Add(hardSoftScore)
				}
			}
		}
	}
	if totalScore == nil {
		return NewHardSoftScore(0, 0, 0)
	}

	return totalScore
}

func (c *IncrementalScoreCalculator) createSubSolution() *solution.SubSolution {
	return solution.NewSubSolution(c.solution, c.dirtyEntities)
}

func (c *IncrementalScoreCalculator) clearDirtyFlags() {
	c.dirtyEntities = make(map[api.IPlanningEntity]struct{})
	c.dirtyVars = make(map[api.IPlanningVariable]struct{})
}
