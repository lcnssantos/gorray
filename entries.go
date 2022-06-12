package goarray

func (a *Array[T]) Entries() []T {
	return a.data
}
