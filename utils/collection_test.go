package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type v struct {
	name  string
	value string
}

type v1 struct {
	name  string
	value string
	id    string
}

func TestCollection(t *testing.T) {
	data := []interface{}{v{"a", "va"}, v{"a1", "va"}, v{"b", "vb"}, v{"b", "vb1"}}
	a := Filter(data, func(i interface{}) bool { return i.(v).name == "b" })
	fmt.Println(a)
	assert.Equal(t, a, []interface{}{v{"b", "vb"}, v{"b", "vb1"}}, "wrong result")
	b := Filter(data, func(i interface{}) bool { return i.(v).value == "va" })
	fmt.Println(b)
	assert.Equal(t, b, []interface{}{v{"a", "va"}, v{"a1", "va"}}, "wrong result")
	c := Map(data, func(i interface{}) interface{} {
		return i.(v).value
	})
	fmt.Println(c)
	assert.Equal(t, c, []interface{}{"va", "va", "vb", "vb1"})
	v1Data := []interface{}{v1{"a", "va", "ka"}, v1{"b", "vb", "kb"}, v1{"c", "vc", "kc"}}
	d := Map(v1Data, func(i interface{}) interface{} {
		return v{i.(v1).name, i.(v1).id}
	})
	fmt.Println(d)
	assert.Equal(t, d, []interface{}{v{"a", "ka"}, v{"b", "kb"}, v{"c", "kc"}})

	e := First(data, func(item interface{}) bool {
		return item.(v).value == "va"
	})
	fmt.Println(e)
	assert.Equal(t, v{"a", "va"}, e, "wrong result")

	f := First(data, func(item interface{}) bool { return item.(v).name == "s" })
	fmt.Println(f)
	assert.Nil(t, f, "weired, must be Nil")

	dataUnify := []interface{}{v{"a", "b"}, v{"a", "b"}, v{"a", "c"}}
	g := Unify(dataUnify)
	fmt.Println(g)
	assert.Equal(t, []interface{}{v{"a", "b"}, v{"a", "c"}}, g, "Not match")

	result := Each(dataUnify, func(arg2 interface{}) interface{} {
		return v{"haha", arg2.(v).value}
	})
	fmt.Println(result)
	assert.Equal(t, []interface{}{v{"haha", "b"}, v{"haha", "b"}, v{"haha", "c"}}, result, "wrong")

	unSorted := []interface{}{v{"c", "2"}, v{"c", "1"}, v{"b", "b"}, v{"d", "d"}}
	target := []interface{}{v{"b", "b"}, v{"c", "1"}, v{"c", "2"}, v{"d", "d"}}
	sorted := Sort(unSorted, func(arg1, arg2 interface{}) bool {
		switch {
		case arg1.(v).name < arg2.(v).name:
			return true
		case arg1.(v).name > arg2.(v).name:
			return false
		default:
			return arg1.(v).value < arg2.(v).value
		}
	})
	fmt.Println("sorted: ", sorted)
	assert.Equal(t, target, sorted, "not sort properly")
}
