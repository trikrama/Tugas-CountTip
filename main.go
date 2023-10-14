package main

import (
	"fmt"
)

func countTip(bill int) string {
	var tip float32
	if bill >= 50 && bill <= 300 {
		tip = float32(bill) * 0.15
	} else {
		tip = float32(bill) * 0.2
	}
	result := float32(bill) + tip
	if tip == float32(int(tip)) {
		return fmt.Sprintf("Tagihannya %d, tipnya %.f, dan total nilainya %.f", bill, tip, result)
	}
	return fmt.Sprintf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f", bill, tip, result)
}

func main() {
	var bill int
	fmt.Print("Masukkan jumlah tagihan: ")
	fmt.Scan(&bill)
	fmt.Println(countTip(bill))
}
