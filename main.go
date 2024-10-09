package main

import (
	"fmt"
	"math"
	"sort"
)

// Fungsi Penjumlahan menggunakan variadic function
func Penjumlahan(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}

// Fungsi Rata-rata menggunakan variadic function
func RataRata(nums ...float64) float64 {
	total := Penjumlahan(nums...)
	return total / float64(len(nums))
}

// Fungsi Median
func Median(nums []float64) float64 {
	sort.Float64s(nums)
	n := len(nums)

	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	}
	return nums[n/2]
}

// Fungsi Luas Lingkaran
func LuasLingkaran(radius float64) float64 {
	return math.Pi * radius * radius
}

// Fungsi Luas Trapesium
func LuasTrapesium(a, b, h float64) float64 {
	return ((a + b) / 2) * h
}

func main() {
	// Contoh penggunaan
	fmt.Println("Penjumlahan:", Penjumlahan(1, 2, 3, 4, 5))                // Output: 15
	fmt.Println("Rata-rata:", RataRata(1, 2, 3, 4, 5))                     // Output: 3
	fmt.Println("Median:", Median([]float64{1, 2, 3, 4, 5}))               // Output: 3
	fmt.Println("Luas Lingkaran (radius 5):", LuasLingkaran(5))            // Output: 78.53981633974483
	fmt.Println("Luas Trapesium (a=5, b=7, h=4):", LuasTrapesium(5, 7, 4)) // Output: 24
}
