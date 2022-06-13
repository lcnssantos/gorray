package goarray

func (a *Array[T]) Push(element T) {
	a.data = append(a.data, element)
}
