package api

// 分数指导器接口 负责计算和缓存分数
type IScoreDirector interface {
	// Calculate 计算完整解决方案的得分
	Calculate(solution ISolution) IScore
	// BeforeVariableChanged 变量改变前的回调
	BeforeVariableChanged(planningVariable IPlanningVariable)
	// AfterVariableChanged 变量改变后的回调
	AfterVariableChanged(planningVariable IPlanningVariable)
	// GetWorkingSolution 获取当前工作解决方案
	GetWorkingSolution() ISolution
	// SetWorkingSolution 设置当前工作解决方案
	SetWorkingSolution(solution ISolution)
}
