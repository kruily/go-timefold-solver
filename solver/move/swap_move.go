package move

import "github.com/kruily/go-timefold-solver/solver/api"

type SwapMove struct {
	entity1       api.IPlanningEntity
	entity2       api.IPlanningEntity
	variable1     api.IPlanningVariable
	variable2     api.IPlanningVariable
	scoreDirector api.IScoreDirector
}

func NewSwapMove(e1, e2 api.IPlanningEntity, v1, v2 api.IPlanningVariable, scoreDirector api.IScoreDirector) *SwapMove {
	return &SwapMove{
		entity1:       e1,
		entity2:       e2,
		variable1:     v1,
		variable2:     v2,
		scoreDirector: scoreDirector,
	}
}

func (m *SwapMove) Execute(workingSolution api.ISolution) {
	oldValue1 := m.variable1.GetValue()
	oldValue2 := m.variable2.GetValue()
	m.scoreDirector.BeforeVariableChanged(m.variable1)
	m.scoreDirector.BeforeVariableChanged(m.variable2)
	m.variable1.SetValue(oldValue2)
	m.variable2.SetValue(oldValue1)
	m.scoreDirector.AfterVariableChanged(m.variable1)
	m.scoreDirector.AfterVariableChanged(m.variable2)
}

func (m *SwapMove) Undo(workingSolution api.ISolution) {
	m.Execute(workingSolution)
}

func (m *SwapMove) Accept(scoreDirector api.IScoreDirector) bool {
	return true
}
