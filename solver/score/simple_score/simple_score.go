package score

import (
	"fmt"
	"math"

	"github.com/kruily/go-timefold-solver/solver/api"
)

type SimpleScore struct {
	score     int
	initScore int
}

func NewSimpleScore(initScore, score int) *SimpleScore {
	return &SimpleScore{score: score, initScore: initScore}
}

var (
	ZERO      = NewSimpleScore(0, 0)
	ONE       = NewSimpleScore(0, 1)
	MINUS_ONE = NewSimpleScore(0, -1)
)

// 解析分数
func parseScore(score string) *SimpleScore {
	return nil
}

// 根据分数值创建分数
func of(score int) *SimpleScore {
	switch score {
	case 0:
		return ZERO
	case 1:
		return ONE
	case -1:
		return MINUS_ONE
	default:
		return NewSimpleScore(score, 0)
	}
}

// 根据初始分数和分数值创建分数
func ofUninitialized(initScore, score int) *SimpleScore {
	if initScore == 0 {
		return of(score)
	}
	return &SimpleScore{score: score, initScore: initScore}
}

func (s *SimpleScore) InitScore() int {
	return s.initScore
}

// 分数值
func (s *SimpleScore) Score() int {
	return s.score
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

func (s *SimpleScore) WithInitScore(score int) api.IScore {
	s.initScore = score
	return s
}

func (s *SimpleScore) Add(score api.IScore) api.IScore {
	return ofUninitialized(
		s.initScore+score.InitScore(),
		s.score+score.(*SimpleScore).Score(),
	)
}

func (s *SimpleScore) Subtract(score api.IScore) api.IScore {
	return ofUninitialized(
		s.initScore-score.InitScore(),
		s.score-score.(*SimpleScore).Score(),
	)
}

func (s *SimpleScore) Multiply(multiplicand float64) api.IScore {
	return ofUninitialized(
		s.initScore*int(multiplicand),
		s.score*int(multiplicand),
	)
}

func (s *SimpleScore) Divide(divisor float64) api.IScore {
	return ofUninitialized(
		int(s.initScore/int(divisor)),
		int(s.score/int(divisor)),
	)
}

func (s *SimpleScore) Power(exponent float64) api.IScore {
	return ofUninitialized(
		int(math.Pow(float64(s.initScore), exponent)),
		int(math.Pow(float64(s.score), exponent)),
	)
}

func (s *SimpleScore) Negate() api.IScore {
	return ofUninitialized(
		-s.initScore,
		-s.score,
	)
}

func (s *SimpleScore) Abs() api.IScore {
	return ofUninitialized(
		int(math.Abs(float64(s.initScore))),
		int(math.Abs(float64(s.score))),
	)
}

func (s *SimpleScore) Zero() api.IScore {
	return ZERO
}

func (s *SimpleScore) IsZero() bool {
	return s.score == 0
}

func (s *SimpleScore) IsSolutionInitailized() bool {
	return s.initScore != 0
}

func (s *SimpleScore) ToLevelNumbers() []int {
	return []int{s.initScore, s.score}
}

func (s *SimpleScore) ToLevelDoubles() []float64 {
	return []float64{float64(s.initScore), float64(s.score)}
}

func (s *SimpleScore) ToShortString() string {
	return fmt.Sprintf("%d", s.score)
}
