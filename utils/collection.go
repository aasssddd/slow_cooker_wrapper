package utils

import "math/rand"

// Filter : Returns a new slice containing all data in the
// slice that satisfy the predicate `f`.
func Filter(vs []interface{}, f func(interface{}) bool) []interface{} {
	vsf := make([]interface{}, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Each : loop execution
func Each(vs []interface{}, do func(interface{}) interface{}) []interface{} {
	vsf := make([]interface{}, 0)
	for i := 0; i < len(vs); i++ {
		newVal := do(vs[i])
		vsf = append(vsf, newVal)
	}
	return vsf
}

// Map : Return a new slice with projected value defined by `f`
func Map(vs []interface{}, mapped func(interface{}) interface{}) []interface{} {
	vsf := make([]interface{}, 0)
	for _, v := range vs {
		vsf = append(vsf, mapped(v))
	}
	return vsf
}

// First : Fetch first item if `condition` match
func First(vs []interface{}, condition func(interface{}) bool) interface{} {
	for _, v := range vs {
		if condition(v) {
			return v
		}
	}
	return nil
}

// Unify : remove duplicate items in slice
func Unify(vs []interface{}) []interface{} {
	vsf := make([]interface{}, 0)
	for _, v := range vs {
		if First(vsf, func(item interface{}) bool { return item == v }) == nil {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Sort : Implementation of Quick sort
func Sort(vs []interface{}, less func(arg1, arg2 interface{}) bool) []interface{} {
	if len(vs) < 2 {
		return vs
	}
	left, right := 0, len(vs)-1
	pivotIndex := rand.Int() % len(vs)

	vs[pivotIndex], vs[right] = vs[right], vs[pivotIndex]
	for i := range vs {
		if less(vs[i], vs[right]) {
			vs[i], vs[left] = vs[left], vs[i]
			left++
		}
	}
	vs[left], vs[right] = vs[right], vs[left]

	Sort(vs[:left], less)
	Sort(vs[left+1:], less)
	return vs
}
