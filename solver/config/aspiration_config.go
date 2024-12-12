package config

import "time"

type AspirationConfig struct {
	Criteria     []AspirationCriteria
	TimeLimit    time.Duration
	MaxFrequency int
}

func NewAspirationConfig(criteria []AspirationCriteria, timeLimit time.Duration, maxFrequency int) *AspirationConfig {
	return &AspirationConfig{
		Criteria:     criteria,
		TimeLimit:    timeLimit,
		MaxFrequency: maxFrequency,
	}
}
