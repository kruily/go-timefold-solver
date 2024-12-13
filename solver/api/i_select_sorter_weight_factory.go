package api

type ISelectionSorterWeightFactory interface {
	CreateSorterWeight(solution ISolution, entity IPlanningEntity) IComparable[any]
}
