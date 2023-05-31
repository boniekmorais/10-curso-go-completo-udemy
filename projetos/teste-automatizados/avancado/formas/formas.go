package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	raio float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func CalcularArea(f Forma) {
	fmt.Printf("A área da forma é: %0.2f\n\n", f.Area())
}

func ExemploInterfaces() {
	fmt.Println("Exemplo de interfaces")

	r := Retangulo{15, 10}
	CalcularArea(r)

	c := Circulo{20}
	CalcularArea(c)

}
