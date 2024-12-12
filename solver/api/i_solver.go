package api

// 规划求解器接口
type ISolver interface {
	// 求解问题
	Solve(problem ISolution) (ISolution, error)
	// 停止求解
	Stop()
	// 是否终止
	IsTerminated() bool
	// 获取最佳解决方案
	GetBestSolution() ISolution
}
