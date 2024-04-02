package main

func main() {
	// x, y := 10, 2
	// a := 5
	// fmt.Printf("x + y = %d \n", x+y)
	// fmt.Printf("x - y = %d \n", x-y)
	// fmt.Printf("x * y = %d \n", x*y)
	// fmt.Printf("x / y = %d \n", x/y)

	// fmt.Printf("x mod y = %d\n", 20%y)

	// a++
	// fmt.Printf("x++ = %d\n", a)

	// a--
	// fmt.Printf("x++ = %d\n", a)

	// var idade int
	// idade = 18

	// idade = idade + 10
	// idade += 10

	// idade = idade - 10
	// idade -= 10

	// idade = idade / 10
	// idade /= 10

	// idade = idade * 10
	// idade *= 10

	// idade = idade % 10
	// idade %= 10

	// x, y, z := 10, 2, 30

	// fmt.Println(x == y)
	// fmt.Println(x != y)
	// fmt.Println(x < y)
	// fmt.Println(x > y)
	// fmt.Println(x >= y)

	// fmt.Println(x <= y)

	// fmt.Println(x < y && y > z)
	// fmt.Println(x < y || y > z)
	// fmt.Println(x < y || y > z)
	// fmt.Println(!(x < y || y > z))

	// var a uint = 60 /* 60 = 0011 1100 */
	// var b uint = 13 /* 13 = 0000 1101 */
	// var c uint = 0

	// aBin := fmt.Sprintf("%b", a)
	// bBin := fmt.Sprintf("%b", b)
	// cBin := fmt.Sprintf("%b", c)

	// fmt.Printf("variavel A é Inicialmente :  %s\n", aBin)
	// fmt.Printf("variavel A é Inicialmente :  %s\n", bBin)
	// fmt.Printf("variavel A é Inicialmente :  %s\n", cBin)

	// var a uint = 60 /* 60 = 0011 1100 */
	// var b uint = 13 /* 13 = 0000 1101 */
	// var c uint = 0

	// c = a & b
	// bin := fmt.Sprintf("%b", c)
	// fmt.Printf("Conjunção - O valor de c é %d e seu binário é  %s\n", c, bin)

	// c = a | b
	// bin := fmt.Sprintf("%b", c)
	// fmt.Printf("Disjunção - O valor de c é %d e seu binário é  %s\n", c, bin)

	// c = a ^ b
	// bin := fmt.Sprintf("%b", c)
	// fmt.Printf("Conjunção - O valor de c é %d e seu binário é  %s\n", c, bin)

	// c = a << 2
	// bin := fmt.Sprintf("%b", c)
	// fmt.Printf("Deslocamento para esquerda - O valor de c é %d e seu binário é  %s\n", c, bin)

	// c = a >> 2
	// bin := fmt.Sprintf("%b", c)
	// fmt.Printf("Deslocamento para esquerda - O valor de c é %d e seu binário é  %s\n", c, bin)

	// var name string
	// var idade int

	// name = "joão"
	// idade = 32

	// fmt.Println(&name)
	// fmt.Println(&idade)

	// x := 32
	// y := x
	// c := &x

	// fmt.Println("endereço de x : ", &x)
	// fmt.Println("endereço de y : ", &y)
	// x = 10

	// fmt.Println("\n\n VALOR de x : ", x)
	// fmt.Println("VALOR de y : ", y)
	// fmt.Println("VALOR de C : ", *c)
	// *c = 20

	// fmt.Println("VALOR de C : ", *c)
	// fmt.Println("VALOR de x : ", x)
	// fmt.Println("VALOR de y : ", y)

	// fmt.Println("VALOR de y : ", &y)
	// fmt.Println("VALOR de y : ", *c)

	// var alunos [2]string

	// fmt.Println(alunos)

	// alunos[0] = "Luan"
	// alunos[1] = "Natan"

	// fmt.Println(alunos)
	// fmt.Println(alunos[0])

	// //slice

	// var s = []string{"Natan", "Lucas", "Luan"}
	// fmt.Println(s)

	// m := make([]int, 10)
	// fmt.Println(m)

	//range

	// primes := []int{1, 2, 3, 4, 5}
	// fmt.Println(primes[1:4])

	// primes = append(primes, 10, 15, 16)
	// fmt.Println(primes)

	myMap := map[string]int{
		"primeira-chave": 25,
	}

	myMap["segunda-chave"] = 15
	myMap["primeira-chave"] = 35

	delete(myMap, "primeira-chave")

}
