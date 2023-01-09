package interfaces

import (
	"fmt"
	"math"
)

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func calcularArea(f forma) {
	fmt.Printf("A área da forma é: %0.2f\n\n", f.area())
}

func ExemploInterfaces() {
	fmt.Println("Exemplo de interfaces")

	r := retangulo{15, 10}
	calcularArea(r)

	c := circulo{20}
	calcularArea(c)

}
