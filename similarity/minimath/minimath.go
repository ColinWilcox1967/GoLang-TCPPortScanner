package minimath

//Max2 returns maximum of two ints
func Max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//Min2 : returns minimum of two ints
func Min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//Min3 Returns smallest of three integers
func Min3(a, b, c int) int {
	if a < b {
		return Min2(a, c)
	}

	return Min2(b, c)
}

//Floor returns largest integer less than value provided
func Floor(f float64) int {
	return int(f)
}

//Ceil returns largets integer above specified value
func Ceil(f float64) int {

	// no fractional part of float

	decPart := int(f)
	if f-float64(decPart) == 0.0 {
		return int(f)
	}

	return int(f + 1)
}

// Max3 returns largest of three integers
func Max3(a, b, c int) int {
	if a > b {
		return Max2(a, c)
	}
	return Max2(b, c)
}

//Absolute Calulate modular value of an integer
func Absolute(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

//AbsoluteF function for real numbers
func AbsoluteF(f float64) float64 {
	if f < 0.0 {
		return -f
	}

	return f
}
