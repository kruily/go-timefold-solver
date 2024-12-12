package api

// 值范围迭代器接口
type IValueRangeIterator interface {
	// HasNext 是否还有下一个值
	HasNext() bool
	// Next 获取下一个值
	Next() interface{}
}
