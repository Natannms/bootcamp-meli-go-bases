package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
	cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
	- Arquivo
	- Nome e sobrenome
	- RG
	- Número de telefone
	- Endereço

	● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
	da cobrança das despesas restantes. Desenvolva e implemente uma função para
	gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
	Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
	interrompa a execução e aborte

	2

	● Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
	fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
	código, implemente a função para ler um arquivo chamado “customers.txt” (como no
	exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
	retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
	agora. Esse erro deve:

	1. Gerar um panic;
	2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
	está danificado”, e continuar com a execução do programa normalmente.

	Requisitos gerais:
	● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
	● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
	um valor de erro (por exemplo, aqueles que tentam ler arquivos).
	Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
	que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
	o caso de erro retornado).
*/

type Customer struct {
	Arquivo   int
	Nome      string
	Sobrenome string
	RG        string
	Telefone  string
	Endereco  string
}

func readCustomers() ([]Customer, error) {
	file, err := os.Open("customers.txt")
	if err != nil {
		return nil, fmt.Errorf("erro: o arquivo indicado não foi encontrado ou está danificado")
	}
	defer file.Close()

	var customers []Customer
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 5 {
			return nil, fmt.Errorf("erro: o arquivo indicado não está no formato correto")
		}

		arquivo, err := strconv.Atoi(fields[0])

		if err != nil {
			return nil, err
		}

		customer := Customer{
			Arquivo:  arquivo,
			Nome:     fields[1],
			RG:       fields[2],
			Telefone: fields[3],
			Endereco: fields[4],
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func saveCustomers(customers []Customer) error {
	file, err := os.Create("customers.txt")
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo")
	}
	defer file.Close()

	for _, customer := range customers {
		_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", customer.Arquivo, customer.Nome, customer.RG, customer.Telefone, customer.Endereco)
		if err != nil {
			return fmt.Errorf("erro ao escrever no arquivo")
		}
	}

	return nil
}

func generateID() (int, error) {
	customers, err := readCustomers()
	if err != nil {
		return 0, err
	}

	return len(customers) + 1, nil
}

func addCustomer(*Customer) error {
	return nil
}

func main() {
	id, err := generateID()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat("customers.txt")
	if err != nil {
		panic(fmt.Errorf("erro: o arquivo indicado não foi encontrado ou está danificado"))
	}

	// Ler os customers existentes do arquivo
	customers, err := readCustomers()
	if err != nil {
		panic(err)
	}

	// criando um novo cliente
	novoCliente := Customer{
		Arquivo:   id,
		Nome:      "João",
		Sobrenome: "Silva",
		RG:        "123456789",
		Telefone:  "(11) 99999-9999",
		Endereco:  "Rua A, 123",
	}

	customers = append(customers, novoCliente)

	// Salvar a lista atualizada no arquivo
	err = saveCustomers(customers)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cliente adicionado com sucesso!")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
}
