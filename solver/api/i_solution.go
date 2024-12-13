package api

// ISolution 解决方案接口
// 实现此接口的结构体可以被solver使用
type ISolution interface {
	// GetScore 获取当前解决方案的评分
	GetScore() IScore
	// SetScore 设置解决方案的评分
	SetScore(score IScore)
	// GetPlanningEntities 获取解决方案中的规划实体
	GetPlanningEntities() []IPlanningEntity
	// SetPlanningEntities 设置解决方案中的规划实体
	SetPlanningEntities(entities []IPlanningEntity)
	// GetProblemFacts 获取问题事实（不可变的问题数据）
	GetProblemFacts() []interface{}
	// SetProblemFacts 设置问题事实（不可变的问题数据）
	SetProblemFacts(facts []interface{})
}
