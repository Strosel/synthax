package synthax

import "reflect"

func Range(a, b, step int) []int {
	out := []int{}
	for i := a; i < b; i += step {
		out = append(out, i)
	}
	return out
}

func RangeF(a, b, step float64) []float64 {
	out := []float64{}
	for i := a; i < b; i += step {
		out = append(out, i)
	}
	return out
}

func Map(set interface{}, fn func(x interface{}) interface{}) []interface{} {
	setVal := reflect.ValueOf(set)
	if setVal.Kind() != reflect.Slice {
		return nil
	}

	out := []interface{}{}

	for i := 0; i < setVal.Len(); i++ {
		var x interface{} = setVal.Index(i)
		if fn != nil {
			x = fn(x)
		}
		out = append(out, x)
	}

	return out
}

func Filter(set interface{}, test func(x interface{}) bool) []interface{} {
	setVal := reflect.ValueOf(set)
	if setVal.Kind() != reflect.Slice {
		return nil
	}

	out := []interface{}{}

	for i := 0; i < setVal.Len(); i++ {
		var x interface{} = setVal.Index(i)
		if test == nil || test(x) {
			out = append(out, x)
		}
	}

	return out
}

func ListComp(fn func(x interface{}) interface{}, set interface{}, test func(x interface{}) bool) []interface{} {
	return Map(Filter(set, test), fn)
}
