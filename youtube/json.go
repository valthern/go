package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string `json:"title"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de NodeJS", 40},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8000", nil)
}
