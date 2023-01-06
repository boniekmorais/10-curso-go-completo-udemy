package arraysslices

import (
	"fmt"
	"reflect"
)

func ArraysSlices() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string

	array1[0] = "texto1"
	array1[2] = "texto3"

	fmt.Println(array1)

	slice1 := []string{}

	slice1 = append(slice1, "slice1")
	slice1 = append(slice1, "slice2")

	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	array2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(array2))

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array2))

	// Arrays internos

	fmt.Println("---------------------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade

}
