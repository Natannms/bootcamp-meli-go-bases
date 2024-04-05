package main

import "fmt"

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, " faz woof woof")
}

func main() {
	s := &Dog{"Marley"}
	s = nil
	s.WoofWoof()
}
