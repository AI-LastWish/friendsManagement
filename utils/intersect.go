package utils

import (
	"reflect"
	"sort"
)

// Simple has complexity: O(n^2)
func SimpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

// Sorted has complexity: O(n * log(n)), a needs to be sorted
func SortedGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool {
			return b[i] == v
		})
		if idx < len(b) && b[idx] == v {
			set = append(set, v)
		}
	}

	return set
}

// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

// Deprecated: Use SimpleGeneric instead. Complexity same as SimpleGeneric.
func Simple(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}

	return set
}

// Deprecated: Use SortedGeneric instead. Complexity same as SortedGeneric.
func Sorted(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		idx := sort.Search(bv.Len(), func(i int) bool {
			return bv.Index(i).Interface() == el
		})
		if idx < bv.Len() && bv.Index(idx).Interface() == el {
			set = append(set, el)
		}
	}

	return set
}

// Deprecated: Use HashGeneric instead. Complexity same as HashGeneric.
func Hash(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]struct{})
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = struct{}{}
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

// Deprecated: Used by Simple which uses reflection.
func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

func AppendWithoutDuplicate(arr []string, text []string) []string {
	textMap := map[string]int{}
	res := make([]string, 0)

	for _, n := range arr {
		if _, ok := textMap[n]; !ok {
			res = append(res, n)
			textMap[n] = 1
		}
	}

	for _, item := range text {
		if _, ok := textMap[item]; !ok {
			res = append(res, item)
		}
	}
	return res
}

func FindMissing(a []string, b []string) []string {
	textMap := map[string]int{}
	res := make([]string, 0)

	for _, n := range b {
		if _, ok := textMap[n]; !ok {
			textMap[n] = 1
		}
	}

	for _, n := range a {
		if _, ok := textMap[n]; !ok {
			res = append(res, n)
		}
	}

	return res
}
