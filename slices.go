package main

import "fmt"

func main()  {
	arreglo := [7]string{"this","is","the","tutorial","of","Go","language"}
	fmt.Println("Arreglo: ",arreglo)
	mySlice := arreglo[1:6]
	fmt.Println("Slice: ",mySlice)
	fmt.Printf("Longitud del slice: %d\n",len(mySlice))
	fmt.Printf("Capacidad del slice: %d\n",cap(mySlice))
	
	slice := make([]string,3,5)
	fmt.Println("\n",slice)
	fmt.Printf("Longitud del slice: %d\n",len(slice))
	fmt.Printf("Capacidad del slice: %d\n",cap(slice))
}
