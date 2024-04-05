package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
	x   string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Ocorreu um erro : %s, %s", e.msg, e.x)
}

func main() {
	e := &myError{"not founded", "404"}

	var err *myError

	isMyError := errors.As(e, &err)

	fmt.Println(isMyError)

}
