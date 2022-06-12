package goarray

func (a *Array[T]) Find(cb func(element T) bool) *T {
	for _, item := range a.data {
		if cb(item) {
			return &item
		}
	}
	return nil
}
