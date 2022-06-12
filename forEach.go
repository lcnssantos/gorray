package goarray

func (a *Array[T]) ForEach(cb func(element T)) {
	for _, item := range a.data {
		cb(item)
	}
}
