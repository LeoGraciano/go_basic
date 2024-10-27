package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

func (r rectangle) zone() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) zone() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type form interface {
	zone() float64
}

func writeZone(f form) {
	fmt.Printf("Area: %.2f\n", f.zone())
}

func generic(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	fmt.Println("Interfaces")
	r := rectangle{5, 3}
	writeZone(r)

	c := circle{7}
	writeZone(c)

	fmt.Println("----------------------")
	println("Tipos Genéricos")
	generic("String")
	generic(10)
	generic(true)

}
