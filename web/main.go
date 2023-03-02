package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor del Header")
		// w.Header().Add("Nombrecito","Valor del Headercito")
		// fmt.Fprintf(w, "Hola mundo!\n")
		// http.Redirect(w,r,"/dos", http.StatusMovedPermanently)
		http.Redirect(w, r, "https://www.google.com.mx", http.StatusMovedPermanently)
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hola mundo, dos!\n")
		http.NotFound(w, r)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
