package config

const (
	// 局部搜索类型
	LocalSearchTypeSimulatedAnnealing = "SIMULATED_ANNEALING"
	LocalSearchTypeTabuSearch         = "TABU_SEARCH"
	LocalSearchTypeLateAcceptance     = "LATE_ACCEPTANCE"
)

type LocalSearchConfig struct {
	// 局部搜索类型
	Type string // "SIMULATED_ANNEALING", "TABU_SEARCH", "LATE_ACCEPTANCE"
	// 接受类型
	AcceptorType string
	// 禁忌步长
	TabuSize int
	// 模拟退火初始温度
	InitialTemperature float64
	// 模拟退火冷却率
	CoolingRate float64
	// 禁忌搜索配置
	TabuSearchConfig TabuSearchConfig
}
