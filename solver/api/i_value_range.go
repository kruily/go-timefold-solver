package api

// 值范围接口
type IValueRange interface {
	// CreateIterator 创建一个迭代器来遍历可能的值
	CreateIterator() IValueRangeIterator
}
