package api

// 规划实体接口
type IPlanningEntity interface {
	// PlanningFilter 获取实体的规划过滤器
	PlanningFilter()
}

// 规划实体注解
type PlanningEntity interface {
	// 获取实体的规划过滤器
	PinningFilter() IPinningFilter
	// 获取空的规划过滤器
	NullPinningFilter() NullPinningFilter
	// 获取实体的难度比较器
	DifficultyComparatorClass() IComparator[IPlanningEntity]
	// 获取空的难度比较器
	NullDifficultyComparator() NullDifficultyComparator
	// 获取实体的难度权重工厂
	DifficultyWeightFactoryClass() ISelectionSorterWeightFactory
}

// 空的规划过滤器
type NullPinningFilter interface {
	IPinningFilter
}

// 空的难度比较器
type NullDifficultyComparator interface {
	IComparator[IPlanningEntity]
}

// 空的难度权重工厂
type NullDifficultyWeightFactory interface {
	ISelectionSorterWeightFactory
}
