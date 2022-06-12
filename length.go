package goarray

func (a *Array[T]) Length() int {
	return len(a.data)
}
