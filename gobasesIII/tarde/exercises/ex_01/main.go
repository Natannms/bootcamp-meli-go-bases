package main

import "fmt"

/*
Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
senha
E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha
*/

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Password  string
}

func (u *User) ChangeName(firstName, lastName string) {
	u.FirstName = firstName
	u.LastName = lastName
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}

func (u *User) ChangeEmail(email string) {
	u.Email = email
}

func (u *User) ChangePassword(password string) {
	u.Password = password
}

func main() {
	// Exemplo de criação de um usuário
	user := User{
		FirstName: "Lucas",
		LastName:  "Alves",
		Age:       20,
		Email:     "lucas@mail.com",
		Password:  "123456",
	}

	// Exemplo de mudança de informações de um usuário
	user.ChangeName("Jane", "Doe")
	user.ChangeAge(35)
	user.ChangeEmail("janedoe@example.com")
	user.ChangePassword("654321")

	// Exibindo informações atualizadas do usuário
	fmt.Printf("%+v", user)
}
