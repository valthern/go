package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("El método es: " + r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola Mundo con el método Get\n")
		case "POST":
			fmt.Fprintf(w, "Hola Mundo con el método Post\n")
		case "PUT":
			fmt.Fprintf(w, "Hola Mundo con el método Put\n")
		case "DELETE":
			fmt.Fprintf(w, "Hola Mundo con el método Delete\n")
		default:
			fmt.Fprintf(w, "Método no válido\n", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
