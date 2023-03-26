package util

import "container/list"

func ListToArray[T any](list *list.List) []T {
	array := make([]T, list.Len())
	for cur, i := list.Front(), 0; cur != nil; cur, i = cur.Next(), i+1 {
		array[i] = cur.Value.(T)
	}
	return array
}
