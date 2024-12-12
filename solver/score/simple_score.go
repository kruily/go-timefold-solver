package score

import "github.com/kruily/go-timefold-solver/solver/api"

type SimpleScore struct {
	score     int
	initScore int
}

func NewSimpleScore(score, initScore int) *SimpleScore {
	return &SimpleScore{score: score, initScore: initScore}
}

func (s *SimpleScore) GetScore() int {
	return s.score
}

func (s *SimpleScore) InitScore() int {
	return s.initScore
}

func (s *SimpleScore) IsFeasible() bool {
	return s.score >= 0
}

func (s *SimpleScore) CompareTo(other api.IScore) int {
	otherScore := other.(*SimpleScore)
	if s.initScore != otherScore.initScore {
		return s.initScore - otherScore.initScore
	}
	return s.score - otherScore.score
}
