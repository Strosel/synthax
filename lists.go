package synthax

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

//todo list comprehension
