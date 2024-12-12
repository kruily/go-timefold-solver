package api

// 分数接口
type IScore interface {
	// 初始化分数
	InitScore() int
	// 是否可行
	IsFeasible() bool
	// 比较分数
	CompareTo(other IScore) int
}
