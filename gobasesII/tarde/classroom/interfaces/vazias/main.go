package main

import "fmt"

type ListaHeterogenea struct {
	Data []interface{}
}

func main() {
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "Olá")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}
