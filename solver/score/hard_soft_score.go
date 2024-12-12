package score

import "github.com/kruily/go-timefold-solver/solver/api"

type HardSoftScore struct {
	hardScore int
	softScore int
	initScore int
}

func NewHardSoftScore(hardScore, softScore, initScore int) *HardSoftScore {
	return &HardSoftScore{hardScore: hardScore, softScore: softScore, initScore: initScore}
}

func (h *HardSoftScore) InitScore() int {
	return h.initScore
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

func (h *HardSoftScore) GetHardScore() int {
	return h.hardScore
}

func (h *HardSoftScore) GetSoftScore() int {
	return h.softScore
}

func (h *HardSoftScore) Add(other api.IScore) *HardSoftScore {
	otherScore := other.(*HardSoftScore)
	return NewHardSoftScore(h.hardScore+otherScore.hardScore, h.softScore+otherScore.softScore, h.initScore)
}
