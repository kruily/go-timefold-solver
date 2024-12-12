package tabu

import (
	"github.com/kruily/go-timefold-solver/solver/api"
	"github.com/kruily/go-timefold-solver/solver/config"
)

type TabuSearchAcceptor struct {
	tabuList        *TabuList
	bestScore       api.IScore
	currentScore    api.IScore
	aspiration      *config.AspirationConfig
	iteration       int
	lastImprovement int
	improvementRate float64
}

func NewTabuSearchAcceptor(minSize, maxSize int, aspiration *config.AspirationConfig) *TabuSearchAcceptor {
	if minSize < 0 {
		minSize = 5
	}
	if maxSize < minSize {
		maxSize = minSize * 2
	}
	return &TabuSearchAcceptor{
		tabuList:        NewTabuList(minSize, maxSize, DefaultMoveHash),
		aspiration:      aspiration,
		iteration:       0,
		lastImprovement: 0,
		improvementRate: 0.9,
	}
}

func (t *TabuSearchAcceptor) Accept(move api.IMove, score api.IScore) (bool, error) {
	isTabu, err := t.tabuList.Contains(move)
	if err != nil {
		return false, err
	}
	if t.currentScore != nil {
		improvement := float64(t.currentScore.CompareTo(score))
		t.improvementRate = 0.9*t.improvementRate + 0.1*improvement
		t.tabuList.adjustSize(improvement)
	}
	if !isTabu {
		return true, nil
	}
	if t.aspiration != nil {
		for _, criteria := range t.aspiration.Criteria {
			if t.checkAspiration(criteria, score) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (t *TabuSearchAcceptor) RecordMove(move api.IMove, score api.IScore) error {
	t.iteration++
	if t.currentScore == nil || score.CompareTo(t.currentScore) > 0 {
		t.lastImprovement = t.iteration
	}

	t.currentScore = score
	if t.bestScore == nil || score.CompareTo(t.bestScore) > 0 {
		t.bestScore = score
	}
	return t.tabuList.Add(move, t.iteration, score)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t *TabuSearchAcceptor) checkAspiration(criteria config.AspirationCriteria, score api.IScore) bool {
	switch criteria {
	case config.BEST_SCORE:
		return t.bestScore != nil && score.CompareTo(t.bestScore) >= 0
	case config.IMPROVING:
		return t.currentScore != nil && score.CompareTo(t.currentScore) >= 0
	case config.TIME_BASED:
		return t.iteration-t.lastImprovement >= int(t.aspiration.TimeLimit.Seconds())
	case config.FREQUENCY_BASED:
		return t.iteration-t.lastImprovement >= t.aspiration.MaxFrequency
	default:
		return false
	}
}

func (t *TabuSearchAcceptor) Clear() {
	t.tabuList.Clear()
	t.bestScore = nil
	t.currentScore = nil
	t.iteration = 0
	t.lastImprovement = 0
	t.improvementRate = 0.9
}
