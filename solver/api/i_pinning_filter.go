package api

type IPinningFilter interface {
	Accept(solution ISolution, entity IPlanningEntity) bool
}
