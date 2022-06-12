package goarray

func (a *Array[T]) Every(callback func(element T) bool) bool {
	for _, element := range a.data {
		result := callback(element)

		if !result {
			return false
		}
	}
	return true
}
