package api

// IConstraintFactory 约束工厂接口
// 实现此接口的结构体可以被solver使用
// 约束工厂用于创建约束
type IConstraintFactory interface {
	// Foreach 遍历所有约束
	Foreach(func(constraint IConstraint)) IConstraintFactory
}
