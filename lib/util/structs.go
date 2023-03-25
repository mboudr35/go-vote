package util

import "container/list"

func ListToArray[T any](list *list.List) []T {
	array := make([]T, 0, list.Len())
	for cur := list.Front(); cur != nil; cur = cur.Next() {
		array = append(array, cur.Value.(T))
	}
	return array
}
