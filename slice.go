package goarray

func (a *Array[T]) Slice() []T {
	return a.data
}
