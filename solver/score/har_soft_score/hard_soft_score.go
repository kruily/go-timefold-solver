package score

import (
	"fmt"
	"math"

	"github.com/kruily/go-timefold-solver/solver/api"
)

var (
	ZERO           = NewHardSoftScore(0, 0, 0)
	ONE_SOFT       = NewHardSoftScore(0, 1, 0)
	ONE_HARD       = NewHardSoftScore(1, 0, 0)
	MINUS_ONE_SOFT = NewHardSoftScore(0, -1, 0)
	MINUS_ONE_HARD = NewHardSoftScore(-1, 0, 0)
)

type HardSoftScore struct {
	hardScore int
	softScore int
	initScore int
}

func NewHardSoftScore(initScore, hardScore, softScore int) *HardSoftScore {
	return &HardSoftScore{hardScore: hardScore, softScore: softScore, initScore: initScore}
}

func ParseScore(score string) *HardSoftScore {
	return nil
}

func ofUninitialized(initScore, hardScore, softScore int) *HardSoftScore {
	if initScore == 0 {
		return of(hardScore, softScore)
	}
	return NewHardSoftScore(initScore, hardScore, softScore)
}

func of(hardScore, softScore int) *HardSoftScore {
	if hardScore == 0 {
		if softScore == -1 {
			return MINUS_ONE_SOFT
		} else if softScore == 0 {
			return ZERO
		} else if softScore == 1 {
			return ONE_SOFT
		}
	} else if softScore == 0 {
		if hardScore == -1 {
			return MINUS_ONE_HARD
		} else if hardScore == 1 {
			return ONE_HARD
		}
	}
	return NewHardSoftScore(0, hardScore, 0)
}

func ofHard(hardScore int) *HardSoftScore {
	if hardScore == -1 {
		return MINUS_ONE_HARD
	} else if hardScore == 1 {
		return ONE_HARD
	} else if hardScore == 0 {
		return ZERO
	}
	return NewHardSoftScore(0, hardScore, 0)
}

func ofSoft(softScore int) *HardSoftScore {
	if softScore == -1 {
		return MINUS_ONE_SOFT
	} else if softScore == 1 {
		return ONE_SOFT
	} else if softScore == 0 {
		return ZERO
	}
	return NewHardSoftScore(0, 0, softScore)
}

func (h *HardSoftScore) InitScore() int {
	return h.initScore
}

func (h *HardSoftScore) HardScore() int {
	return h.hardScore
}

func (h *HardSoftScore) SoftScore() int {
	return h.softScore
}

func (h *HardSoftScore) IsFeasible() bool {
	return h.initScore >= 0 && h.hardScore >= 0
}

func (h *HardSoftScore) CompareTo(other api.IScore) int {
	otherScore := other.(*HardSoftScore)
	if h.initScore != otherScore.initScore {
		return h.initScore - otherScore.initScore
	}
	if h.hardScore != otherScore.hardScore {
		return h.hardScore - otherScore.hardScore
	}
	return h.softScore - otherScore.softScore
}

func (h *HardSoftScore) WithInitScore(score int) api.IScore {
	h.initScore = score
	return h
}

func (h *HardSoftScore) Add(other api.IScore) api.IScore {
	return ofUninitialized(
		h.initScore+other.InitScore(),
		h.hardScore+other.(*HardSoftScore).HardScore(),
		h.softScore+other.(*HardSoftScore).SoftScore(),
	)
}

func (h *HardSoftScore) Subtract(other api.IScore) api.IScore {
	return ofUninitialized(
		h.initScore-other.InitScore(),
		h.hardScore-other.(*HardSoftScore).HardScore(),
		h.softScore-other.(*HardSoftScore).SoftScore(),
	)
}

func (h *HardSoftScore) Multiply(multiplicand float64) api.IScore {
	return ofUninitialized(
		h.initScore*int(multiplicand),
		h.hardScore*int(multiplicand),
		h.softScore*int(multiplicand),
	)
}

func (h *HardSoftScore) Divide(divisor float64) api.IScore {
	return ofUninitialized(
		h.initScore/int(divisor),
		h.hardScore/int(divisor),
		h.softScore/int(divisor),
	)
}

func (h *HardSoftScore) Power(exponent float64) api.IScore {
	return ofUninitialized(
		h.initScore*int(exponent),
		h.hardScore*int(exponent),
		h.softScore*int(exponent),
	)
}

func (h *HardSoftScore) Abs() api.IScore {
	return ofUninitialized(
		int(math.Abs(float64(h.initScore))),
		int(math.Abs(float64(h.hardScore))),
		int(math.Abs(float64(h.softScore))),
	)
}

func (h *HardSoftScore) Zero() api.IScore {
	return ZERO
}

func (h *HardSoftScore) IsZero() bool {
	return h.hardScore == 0 && h.softScore == 0
}

func (h *HardSoftScore) Negate() api.IScore {
	return ofUninitialized(
		-h.initScore,
		-h.hardScore,
		-h.softScore,
	)
}

func (h *HardSoftScore) ToLevelNumbers() []int {
	return []int{h.initScore, h.hardScore, h.softScore}
}

func (h *HardSoftScore) ToLevelDoubles() []float64 {
	return []float64{float64(h.initScore), float64(h.hardScore), float64(h.softScore)}
}

func (h *HardSoftScore) IsSolutionInitailized() bool {
	return h.initScore != 0
}

func (h *HardSoftScore) ToShortString() string {
	return fmt.Sprintf("HardSoftScore[initScore=%d, hardScore=%d, softScore=%d]", h.initScore, h.hardScore, h.softScore)
}
