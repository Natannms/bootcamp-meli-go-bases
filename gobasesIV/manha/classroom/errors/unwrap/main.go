package main

import (
	"errors"
	"fmt"
)

type erroTwo struct{}

func (e erroTwo) Error() string {
	return "error two happened"
}

func main() {
	e2 := erroTwo{}
	e1 := fmt.Errorf("e2: %w", e2)

	fmt.Println(errors.Unwrap(e1))
	fmt.Println(errors.Unwrap(e2))
}
