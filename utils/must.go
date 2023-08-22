package utils

// panic if err, else return val
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// similar to must but returns true always if theres no err. For err it just panics
func MustB[T any](val T, err error) bool {
	Must(val, err)
	return true
}
