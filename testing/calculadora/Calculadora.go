package calculadora

import (
	"fmt"
	"sort"
	"errors"
)

func restar(a int, b int) int {
	return a - b
}

func Main() {
	result := restar(5,1)
	fmt.Println(result)
}


func Sort(array []int) []int { 
	sort.Slice(array, func(i,j int) bool { return array[i] < array[j] })
	return array
}

func Dividir(num, den int) (int,error) {
	if (den == 0) {
		return 0, errors.New("denominator must be different than 0")
	}
	return num / den, nil
}