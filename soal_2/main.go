package main

import (
	"fmt"
)

// Struct untuk menyimpan informasi individu
type Person struct {
	Berat float64 // kg
	Tinggi float64 // meter
}

// Method untuk menghitung BMI seseorang
func (p Person) CalculateBMI() float64 {
	return p.Berat / (p.Tinggi * p.Tinggi)
}

func main() {
	// Data uji
	mark := Person{78, 1.69}
	john := Person{92, 1.95}

	// Menghitung BMI untuk Mark dan John
	markBMI := mark.CalculateBMI()
	johnBMI := john.CalculateBMI()

	// Mencetak BMI untuk Mark dan John
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)

	// Membandingkan BMI Mark dan John
	markHigherBMI := markBMI > johnBMI

	// Mencetak hasil perbandingan
	if markHigherBMI {
		fmt.Println("Mark memiliki BMI lebih tinggi dari John.")
	} else {
		fmt.Println("John memiliki BMI lebih tinggi dari Mark.")
	}
}