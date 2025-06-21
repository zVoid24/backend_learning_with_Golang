package mathlib

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a float32, b float32) float32 {
	if b != 0 {
		return a / b
	}
	return b / a
}
