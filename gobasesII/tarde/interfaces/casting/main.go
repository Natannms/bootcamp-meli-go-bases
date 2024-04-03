package main

import (
	"fmt"
)

func main() {
	// o tipo any é um alias para interface, produz o mesmo resultado que interface{}
	// var i any = "hello"
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
