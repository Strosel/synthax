package synthax

func Ternary(test bool, a, b interface{}) interface{} {
	if test {
		return a
	}
	return b
}
