package move

import "github.com/kruily/go-timefold-solver/solver/api"

type ChangeMove struct {
	entity        api.IPlanningEntity
	variable      api.IPlanningVariable
	targetValue   interface{}
	scoreDirector api.IScoreDirector
}

func NewChangeMove(entity api.IPlanningEntity, variable api.IPlanningVariable, targetValue interface{}, scoreDirector api.IScoreDirector) *ChangeMove {
	return &ChangeMove{
		entity:        entity,
		variable:      variable,
		targetValue:   targetValue,
		scoreDirector: scoreDirector,
	}
}

func (m *ChangeMove) Execute(workingSolution api.ISolution) {
	oldValue := m.variable.GetValue()
	m.scoreDirector.BeforeVariableChanged(m.variable)
	m.variable.SetValue(m.targetValue)
	m.scoreDirector.AfterVariableChanged(m.variable)
	m.targetValue = oldValue // 保存旧值，用于撤销
}

func (m *ChangeMove) Undo(workingSolution api.ISolution) {
	m.variable.SetValue(m.targetValue)
}

func (m *ChangeMove) Accept(scoreDirector api.IScoreDirector) bool {
	return true
}
