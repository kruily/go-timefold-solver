package api

// 移动接口
type IMove interface {
	// 执行移动
	Execute(workingSolution ISolution)
	// 撤销移动
	Undo(workingSolution ISolution)
	// 接受移动
	Accept(scoreDirector IScoreDirector) bool
}
