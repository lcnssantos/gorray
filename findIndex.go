package goarray

func (a *Array[T]) FindIndex(cb func(element T) bool) int {
	for index, item := range a.data {
		if cb(item) {
			return index
		}
	}
	return -1
}
