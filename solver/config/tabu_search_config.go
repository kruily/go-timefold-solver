package config

import "time"

// AspirationCriteria 特赦准则类型
type AspirationCriteria int

const (
	NONE            AspirationCriteria = iota
	BEST_SCORE                         // 超过历史最佳分数
	IMPROVING                          // 比当前解更好
	TIME_BASED                         // 基于时间的特赦
	FREQUENCY_BASED                    // 基于频率的特赦
)

// 禁忌搜索配置
type TabuSearchConfig struct {
	// 最小禁忌步长
	MinTabuSize int
	// 最大禁忌步长
	MaxTabuSize int
	// 接受条件
	AspirationCriteria []AspirationCriteria
	// 时间限制
	TimeLimit time.Duration
	// 最大频率
	MaxFrequency int
}
