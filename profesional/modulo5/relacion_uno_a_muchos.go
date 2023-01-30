package main

import "fmt"

type Curso struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Curso  Curso
}

func main() {
	video1 := Video{Titulo: "1.- Introducción"}
	video2 := Video{Titulo: "2.- Instalación"}

	videos := []Video{video1, video2}

	curso := Curso{Titulo: "Curso profesional de Go!", Videos: videos}

	video1.Curso = curso
	video2.Curso = curso

	fmt.Println(video1.Curso.Titulo)
	
	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}
}
