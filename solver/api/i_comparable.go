package api

type IComparable[T any] interface {
	CompareTo(other T) int
}
