package day_13

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

// luas() get from interface hitung
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

// keliling() get from interface hitung
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

// luas() get from interface hitung
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

// keliling() get from interface hitung
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func Interface() {
	var bangunDatar hitung

	bangunDatar = persegi{5.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{21.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

}
