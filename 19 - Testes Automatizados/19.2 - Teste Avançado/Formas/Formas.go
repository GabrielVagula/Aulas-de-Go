package Formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

/* func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
} */

func main() {
	r := Retangulo{10, 15}

	c := Circulo{10}

	EscreverArea(r)
	EscreverArea(c)

}
