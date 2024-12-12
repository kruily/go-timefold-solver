package config

import "time"

const (
	MOVE_SELECTOR_FIRST_FIT = "FIRST_FIT"
	MOVE_SELECTOR_BEST_FIT  = "BEST_FIT"
	MOVE_SELECTOR_CHANGE    = "CHANGE"
	MOVE_SELECTOR_CHAINED   = "CHAINED"
	MOVE_SELECTOR_RANDOM    = "RANDOM"
)

type SolverConfig struct {
	// 求解时间限制（秒）
	TimeLimit int
	// 是否使用并行求解
	Parallel bool
	// 并行线程数
	ParallelThreadCount int
	// 移动选择策略
	MoveSelector string
	// 是否启用局部搜索
	LocalSearch bool
	// 终止配置
	Termination TerminationConfig
	// 构造启发式配置
	ConstructionHeuristic string
	// 局部搜索配置
	LocalSearchConfig LocalSearchConfig
	// 是否启用邻域缓存
	NeighborhoodCaching bool
	// 随机种子
	RandomSeed int64
}

func NewDefalutSolverConfig() *SolverConfig {
	return &SolverConfig{
		TimeLimit:           60,
		Parallel:            true,
		ParallelThreadCount: 4,
		MoveSelector:        MOVE_SELECTOR_BEST_FIT,
		LocalSearch:         true,
		Termination: TerminationConfig{
			TimeLimit:                300,
			UnimprovedStepCountLimit: 100,
		},
		ConstructionHeuristic: "FIRST_FIT_DECREASING",
		LocalSearchConfig: LocalSearchConfig{
			Type:               LocalSearchTypeSimulatedAnnealing,
			AcceptorType:       LocalSearchTypeSimulatedAnnealing,
			InitialTemperature: 1000,
			CoolingRate:        0.99,
		},
		NeighborhoodCaching: true,
		RandomSeed:          time.Now().UnixNano(),
	}
}
