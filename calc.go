package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Tambah(a, b int) int {
	return a + b
}

func Kurang(a, b int) int {
	return a - b
}

func Recover () {
	r:= recover()

	if r != nil {
        fmt.Println(r)
    } else {
        fmt.Println("\n\nApplication running perfectly")
    }
}

func main() {
	defer Recover()
	var bil1, bil2 string

	fmt.Print("\nKalkulator Penjumlahan & Pengurangan 2 Bilangan\n\n")
	fmt.Print("Masukkan Bilangan Pertama : ")
	fmt.Scanln(&bil1)
	fmt.Print("Masukkan Bilangan Kedua : ")
	fmt.Scanln(&bil2)
	fmt.Print("\n\n")

	var num1 int
	num1, err := strconv.Atoi(bil1)

	if err == nil {
		var num2 int
		num2, err := strconv.Atoi(bil2)

		if err == nil {
			fmt.Println("Hasil Penjumlahan ", bil1, " + ", bil2, " = ", Tambah(num1, num2))
			fmt.Println("Hasil Pengurangan ", bil1, " - ", bil2, " = ", Kurang(num1, num2))
		} else {
			// fmt.Println()
			panic(errors.New("Bilangan kedua bukan angka"))
		}
	} else {
		// fmt.Println()
		panic(errors.New("Bilangan pertama bukan angka"))
	}
}