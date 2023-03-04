package ponteiros

import "fmt"

func Ponteiros() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1
	var p *int = &v1
	p2 := &v1

	v1++

	fmt.Println(v1, v2)

	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(p2)
	fmt.Println(*p2)

	fmt.Println(&v1)
	fmt.Println(&v2)

}
