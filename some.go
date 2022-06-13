package goarray

func (a *Array[T]) Some(cb func(element T) bool) bool {
	for _, element := range a.data {
		result := cb(element)

		if result {
			return true
		}
	}
	return false
}
