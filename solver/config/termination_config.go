package config

import "github.com/kruily/go-timefold-solver/solver/api"

type TerminationConfig struct {
	// 时间限制（秒）
	TimeLimit int
	// 未改进步数限制
	UnimprovedStepCountLimit int
	// 最佳分数限制
	BestScoreLimit api.IScore
	// 步数限制
	StepCountLimit int
}
