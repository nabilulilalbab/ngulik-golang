package ngulik

import (
	"fmt"
	"math"
	"testing"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type Lingkaran struct {
	diameter float64
}

func (l Lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l Lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l Lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type Persegi struct {
	sisi float64
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) keliling() float64 {
	return p.sisi * 4
}

func TestKontrakInterface(t *testing.T) {
	var bangunDatar hitung
	bangunDatar = Persegi{
		sisi: 10.0,
	}
	fmt.Println("=========Persegi")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling : ", bangunDatar.keliling())
	bangunDatar = Lingkaran{
		diameter: 14.0,
	}
	fmt.Println("=========Lingkaran")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling : ", bangunDatar.keliling())
	// type assertion untuk menentukan tipe data yang benar bukan casting tipe data
	fmt.Println("jariJari : ", bangunDatar.(Lingkaran).jariJari())
}
