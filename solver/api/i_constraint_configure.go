package api

// 约束配置接口
type IConstraintConfigure interface {
	// 获取约束配置
	GetConstraints() []IConstraint
}
