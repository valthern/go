package main

import (
	"modificadores/CodigoFacilito"
	"fmt"
)

func main() {
	curso := modificadores.Curso{Titulo: "Curso profesional de Go!"}

	fmt.Println(curso.ObtenerTitulo())
}
