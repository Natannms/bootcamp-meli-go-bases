package main

import (
	"errors"
	"fmt"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return fmt.Sprintf("Hello %s ", name), nil
}

func main() {
	name := ""
	msg, err := SayHello(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(msg)
}
