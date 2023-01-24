package main

/*
Un paquete es una carpeta que almacena archivos ".go"
*/

/**
*Encontré una solución para eso. En la carpeta raíz,
tenes que generar un modulo con un nombre con el comando:
go mod init <Nombre del modulo>
, y en el import colocas el "<nombre del modulo>/CodigoFacilito".
*/

import (
	"aplicacion/CodigoFacilito"
	"fmt"
)

func main() {
	curso := CodigoFacilito.Curso{Titulo: "Curso profesional de Go!"}

	// NO ENCUENTRA EL MÉTODO getTitulo()
	fmt.Println(curso.GetTitulo())
}
