package util

import "container/list"

func ListToArray[T any](list *list.List) []T {
	array := make([]T, list.Len())
	for cur, i := list.Front(), 0; cur != nil; cur, i = cur.Next(), i+1 {
		array[i] = cur.Value.(T)
	}
	return array
}

func ArrayDeleteElement[T comparable](array []T, x T) []T {
	var i int
	for i = 0; i < len(array); i++ {
		if array[i] == x {
			break
		}
	}
	return append(array[:i], array[i+1:]...)
}
