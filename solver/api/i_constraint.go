package api

// 约束接口
type IConstraint interface {
	// 获取约束的得分
	GetScore() IScore
	// 检查约束是否满足
	Match(solution ISolution) bool
	// 获取约束权重
	GetWeight() int
}
