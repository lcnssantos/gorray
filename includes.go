package goarray

import "reflect"

func (a *Array[T]) Includes(element T) bool {
	for _, item := range a.data {
		if reflect.DeepEqual(item, element) {
			return true
		}
	}

	return false
}
