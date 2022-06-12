package goarray

func (a *Array[T]) Filter(cb func(element T) bool) Array[T] {
	data := make([]T, 0)

	for _, item := range a.data {
		if cb(item) {
			data = append(data, item)
		}
	}

	return New(&data)
}
