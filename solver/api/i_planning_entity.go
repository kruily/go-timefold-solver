package api

// 规划实体接口
type IPlanningEntity interface {
	// GetPlanningVariables 获取实体的所有规划变量
	GetPlanningVariables() []IPlanningVariable
}
