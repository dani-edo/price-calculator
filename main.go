package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64) // make: create an empty map (len shouldn't be defined when creating the map)

	for _, taxRate := range taxRates {
		taxIndludedPrices := make([]float64, len(prices)) // make: create an empty slice (len should be defined when creating the slice)
		for priceIndex, price := range prices {
			taxIndludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIndludedPrices
	}

	fmt.Println(result)
}
