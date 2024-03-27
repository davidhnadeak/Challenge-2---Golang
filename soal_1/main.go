package main

import (
	"fmt"
)

// Tim struct
type Tim struct {
	nama  string
	skor []int
}

// Method untuk menghitung skor rata-rata tim
func (t Tim) skorRataRata() float64 {
	total := 0
	for _, skor := range t.skor {
		total += skor
	}
	return float64(total) / float64(len(t.skor))
}

// Function untuk menentukan pemenang berdasarkan skor rata-rata
func tentukanPemenang(tim1, tim2 Tim) string {
	skorTim1 := tim1.skorRataRata()
	skorTim2 := tim2.skorRataRata()

	fmt.Printf("Skor %s: %.2f\n", tim1.nama, skorTim1)
	fmt.Printf("Skor %s: %.2f\n", tim2.nama, skorTim2)

	// Cek skor minimum 100
	if skorTim1 >= 100 && skorTim2 >= 100 {
		if skorTim1 > skorTim2 {
			return fmt.Sprintf("Pemenang: %s", tim1.nama)
		} else if skorTim2 > skorTim1 {
			return fmt.Sprintf("Pemenang: %s", tim2.nama)
		} else {
			return "Hasil seri!"
		}
	} else {
		return "Tidak ada pemenang, skor di bawah 100 poin!"
	}
}

func main() {
	// Data 1
	lumbaLumba := Tim{
		nama:  "Tim Lumba-lumba",
		skor: []int{96, 108, 89},
	}

	koala := Tim{
		nama:  "Tim Koala",
		skor: []int{88, 91, 110},
	}

	fmt.Println("Data 1:")
	fmt.Println(tentukanPemenang(lumbaLumba, koala))

	// Data Bonus 1
	lumbaLumbaBonus1 := Tim{
		nama:  "Tim Lumba-lumba",
		skor: []int{97, 112, 101},
	}

	koalaBonus1 := Tim{
		nama:  "Tim Koala",
		skor: []int{109, 95, 123},
	}

	fmt.Println("\nData Bonus 1:")
	fmt.Println(tentukanPemenang(lumbaLumbaBonus1, koalaBonus1))

	// Data Bonus 2
	lumbaLumbaBonus2 := Tim{
		nama:  "Tim Lumba-lumba",
		skor: []int{97, 112, 101},
	}

	koalaBonus2 := Tim{
		nama:  "Tim Koala",
		skor: []int{109, 95, 106},
	}

	fmt.Println("\nData Bonus 2:")
	fmt.Println(tentukanPemenang(lumbaLumbaBonus2, koalaBonus2))
}