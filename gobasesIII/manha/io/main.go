package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	print("=========================== COPIA DO LEITOR PARA DESTINO ==========================\n")
	// usamos o pacote strings para transformar a string em um io.Reader

	/*
		    formattedString := fmt.Sprintf("some io.Reader stream to be read [bytes]")
			  reader := bytes.NewReader([]byte(formattedString))
	*/

	r := strings.NewReader("Texto")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println(err)
	}

	print("=========================== IMPRESSAO NO TERMINAL COM WriteString ==========================\n")
	io.WriteString(os.Stdout, "Hello world!")

}
