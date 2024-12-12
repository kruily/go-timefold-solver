package solution

import "github.com/kruily/go-timefold-solver/solver/api"

type SubSolution struct {
	originalSolution api.ISolution
	dirtyEntities    map[api.IPlanningEntity]struct{}
}

func NewSubSolution(originalSolution api.ISolution, dirtyEntities map[api.IPlanningEntity]struct{}) *SubSolution {
	return &SubSolution{
		originalSolution: originalSolution,
		dirtyEntities:    dirtyEntities,
	}
}

func (s *SubSolution) GetOriginalSolution() api.ISolution {
	return s.originalSolution
}

func (s *SubSolution) GetDirtyEntities() map[api.IPlanningEntity]struct{} {
	return s.dirtyEntities
}

func (s *SubSolution) GetScore() api.IScore {
	return s.originalSolution.GetScore()
}

func (s *SubSolution) SetScore(score api.IScore) {
	s.originalSolution.SetScore(score)
}

func (s *SubSolution) GetProblemFacts() []interface{} {
	return s.originalSolution.GetProblemFacts()
}
