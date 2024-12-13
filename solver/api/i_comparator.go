package api

type IComparator[T any] interface {
	Compare(a, b T) int
}
