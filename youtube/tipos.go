package main

import(
	"fmt"
	"strconv"
)

func main(){
	edad := "22"
	edad_int,_ := strconv.Atoi(edad)

	fmt.Println("Mi edad es",edad_int)
	// fmt.Println(err)
}
