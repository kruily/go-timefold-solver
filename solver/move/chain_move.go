package move

import "github.com/kruily/go-timefold-solver/solver/api"

type ChainMove struct {
	moveList      []api.IMove
	scoreDirector api.IScoreDirector
}

func NewChainMove(moveList []api.IMove, scoreDirector api.IScoreDirector) *ChainMove {
	return &ChainMove{
		moveList:      moveList,
		scoreDirector: scoreDirector,
	}
}

func (m *ChainMove) Execute(workingSolution api.ISolution) {
	for _, move := range m.moveList {
		move.Execute(workingSolution)
	}
}

func (m *ChainMove) Undo(workingSolution api.ISolution) {
	for i := len(m.moveList) - 1; i >= 0; i-- {
		m.moveList[i].Undo(workingSolution)
	}
}

func (m *ChainMove) Accept(scoreDirector api.IScoreDirector) bool {
	return true
}
