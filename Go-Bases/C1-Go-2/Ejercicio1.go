package main

import (
	"fmt"
)

func main() {
	palabra :=  "Programar"
	length := len(palabra)
	fmt.Println(length, '\n')
	for i := 0; i < length; i++ {
		fmt.Println(palabra[i:i+1]);
	}
}
