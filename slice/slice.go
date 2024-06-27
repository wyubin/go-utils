package slice

import "golang.org/x/exp/constraints"

// return unique slice items
func Uniq[T constraints.Ordered](items []T) []T {
	itemMap := map[T]struct{}{}
	res := []T{}
	for _, _item := range items {
		if _, ok := itemMap[_item]; ok {
			continue
		}
		itemMap[_item] = struct{}{}
		res = append(res, _item)
	}
	return res
}

// first items is subset slice
func Subset[T constraints.Ordered](subset, total []T) bool {
	set := make(map[T]int)
	for _, value := range total {
		set[value] += 1
	}
	for _, value := range subset {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

// get intersection of slices
func Intersect[T constraints.Ordered](a, b []T) []T {
	set := Map(a)
	res := []T{}
	for _, value := range b {
		if _, found := set[value]; !found {
			res = append(res, value)
		}
	}
	return res
}

// get map from slice
func Map[T constraints.Ordered](slice []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, value := range slice {
		res[value] = struct{}{}
	}
	return res
}

// remove elements from slice
func Remove[T constraints.Ordered](slice []T, item T) []T {
	res := []T{}
	for _, value := range slice {
		if value != item {
			res = append(res, value)
		}
	}
	return res
}

func Transpose[T constraints.Ordered](slice [][]T) [][]T {
	res := make([][]T, len(slice[0]))
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[0]); j++ {
			res[j] = append(res[j], slice[i][j])
		}
	}
	return res
}
