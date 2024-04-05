package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Nome string
		}{}
		defer r.Body.Close()
		fmt.Println(body)
	})

	http.ListenAndServe(":8080", nil)
}

// chunks
// bata | tinha | frita | 123 |

// requisições http trabalham com streams, informações chegam através de um "tubo"
// sob demanda
