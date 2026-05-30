package main

import (
	"tax_prices/cmdmanager"
	"tax_prices/prices"
)

func main() {
	taxes := []float64{0, 0.05, 0.1, 0.15}

	for _, taxValue := range taxes {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxValue*100))
		// taxIncluder := prices.New(fm, taxValue)

		cmdm := cmdmanager.New()
		taxIncluder := prices.New(cmdm, taxValue)

		taxIncluder.Process()
	}
}
