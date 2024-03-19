package try

import "fmt"

func Try(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	return fn()
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
