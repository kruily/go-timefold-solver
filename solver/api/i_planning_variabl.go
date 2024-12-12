package api

type IPlanningVariable interface {
	// GetValue 获取变量的当前值
	GetValue() interface{}
	// SetValue 设置变量的值
	SetValue(value interface{})
	// GetValueRange 获取变量的可能值范围
	GetValueRange() IValueRange
}
