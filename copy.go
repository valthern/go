package main

import "fmt"

func main()  {
	/*
		La funcion 'copy' copia el mínimo de elementos en ambos arreglos 
	*/
	slice := []int{1,2,3,4,5,6,7,8,9}
	copia := make([]int,len(slice),cap(slice)*2)

	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)
}