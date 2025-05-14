package main

import "fmt"

func nilaiSaham(hargas []int) int {
	if len(hargas) < 2 {
		return -1
	}

	minHarga := hargas[0]
	maxProfit := 0
	bestBuyPrice := hargas[0]

	for _, harga := range hargas[1:] {
		profit := harga - minHarga

		if profit > maxProfit {
			maxProfit = profit
			bestBuyPrice = minHarga
		}
		if harga < minHarga {
			minHarga = harga
		}
	}

	if maxProfit == 0 {
		return -1
	}
	return bestBuyPrice
}

func main() {
	testCase := [][]int{
		{7, 8, 3, 10, 8},
		{5, 12, 11, 12, 10},
		{7, 18, 27, 10, 29},
		{20, 17, 15, 14, 10},
	}

	for i, test := range testCase {
		fmt.Printf("Soal %d: %v\n", i+1, test)
		fmt.Printf("Output: %d\n\n", nilaiSaham(test))
	}
}
