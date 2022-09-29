package main

import (
	"fmt"
	"errors"
)

type Matrix struct {
	rows int
    cols int
    data [][]float64
}


func (m *Matrix) Set(x, y int, nums ...float64) (err error) {
	m.rows = x
	m.cols = y
	if len(nums) != x*y{
		err = errors.New("El tamaño de la matriz no coincide con la cantidad de datos recibidos.")
		return
	}
	
	m.data = make([][]float64, x)
	for i := 0; i < x; i++ {
		m.data[i] = make([]float64, y)
        for j := 0; j < y; j++ {
            m.data[i][j] = nums[i*y+j]
        }
	}
	return
}

func (m *Matrix) Show() {
	for i := 0; i < m.rows; i++{
        for j := 0; j < m.cols; j++ {
            fmt.Print(m.data[i][j]," ")
        }
		fmt.Println()
	}
}

func ok(err error) (bool){
	if (err!= nil) {
        fmt.Println(err.Error())
		return false
    }
	return true
}


func main() {
	m := Matrix{}
	m.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if ok(m.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)) {
		m.Show()
	}
}
