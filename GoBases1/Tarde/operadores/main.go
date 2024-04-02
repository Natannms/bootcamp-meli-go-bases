package main

import "fmt"

func main() {
	//operadores aritiméticos
	x, y := 10, 2

	fmt.Printf("x + y = %d \n", x+y)
	fmt.Printf("x - y = %d \n", x-y)
	fmt.Printf("x * y = %d \n", x*y)
	fmt.Printf("x / y = %d \n", x/y)
	fmt.Printf("x mod y = %d\n", 20%y)

	incremento := 5
	decremento := 10

	incremento++
	fmt.Printf("x++ = %d\n", incremento)

	decremento--
	fmt.Printf("x++ = %d\n", decremento)

	// operadores de atribuição
	var idade int
	idade = 18

	idade = idade + 10
	idade += 10

	idade = idade - 10
	idade -= 10

	idade = idade / 10
	idade /= 10

	idade = idade * 10
	idade *= 10

	idade = idade % 10
	idade %= 10

	//operadores de comparação
	x, y, z := 10, 2, 30
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	//operadores logicos
	fmt.Println(x < y && y > z)
	fmt.Println(x < y || y > z)
	fmt.Println(x < y || y > z)
	fmt.Println(!(x < y || y > z))

	//Operadores logicos a nivel de bit (bit a bit)
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b
	bin := fmt.Sprintf("%b", c)
	fmt.Printf("Conjunção - O valor de c é %d e seu binário é  %s\n", c, bin)

	c = a | b
	conjucao := fmt.Sprintf("%b", c)
	fmt.Printf("Disjunção - O valor de c é %d e seu binário é  %s\n", c, conjucao)

	c = a ^ b
	disjuncao := fmt.Sprintf("%b", c)
	fmt.Printf("Conjunção - O valor de c é %d e seu binário é  %s\n", c, disjuncao)

	c = a << 2
	deslocamentoEsquerda := fmt.Sprintf("%b", c)
	fmt.Printf("Deslocamento para esquerda - O valor de c é %d e seu binário é  %s\n", c, deslocamentoEsquerda)

	c = a >> 2
	deslocamentoDireta := fmt.Sprintf("%b", c)
	fmt.Printf("Deslocamento para esquerda - O valor de c é %d e seu binário é  %s\n", c, deslocamentoDireta)
}
