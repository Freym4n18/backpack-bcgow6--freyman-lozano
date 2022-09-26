package main

import (
	"fmt"
	"errors"
)

const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
)

var kgPorTipo = map[string]float64{"dog":10,"cat":5,"hamster":0.250,"tarantula":0.150}

func Animal(animal string) (kg float64, err error){
	var exists bool
	_,exists = kgPorTipo[animal]
	if !exists {
		err = errors.New("El animal indicado no esta registrado")
	}
	kg = kgPorTipo[animal]
	return
}

func calculateAmount(animal string, number int) (float64){
	kg, err := Animal(animal)
	if (err != nil) {
		fmt.Println(err)
		return 0
	}
	return kg * float64(number)
}


func main() {
	amount := calculateAmount(dog,5)
	amount += calculateAmount(cat,51)
	amount += calculateAmount(hamster,2)
	fmt.Println(amount)
}
