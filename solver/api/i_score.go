package api

// 分数接口
type IScore interface {
	// 初始化分数s
	InitScore() int
	// 是否可行
	IsFeasible() bool
	// 比较分数
	CompareTo(other IScore) int
	// 初始化分数
	WithInitScore(score int) IScore
	// 添加分数
	Add(score IScore) IScore
	// 减去分数
	Subtract(score IScore) IScore
	// 乘以分数
	Multiply(multiplicand float64) IScore
	// 除以分数
	Divide(divisor float64) IScore
	// 乘方
	Power(exponent float64) IScore
	// 取反
	Negate() IScore
	// 取绝对值
	Abs() IScore
	Zero() IScore
	// 是否为0
	IsZero() bool
	// 转换为数字
	ToLevelNumbers() []int
	// 转换为浮点数
	ToLevelDoubles() []float64
	// 是否已初始化
	IsSolutionInitailized() bool
	// 转换为短字符串
	ToShortString() string
}
