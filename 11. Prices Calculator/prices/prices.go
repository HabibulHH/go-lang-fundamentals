package prices

import (
	"fmt"
	"tax_prices/conversion"
	"tax_prices/iomanager"
)

type TaxIncluder struct {
	IOManager   iomanager.IOManager `json:"-"`
	TaxRate     float64             `json:"tax_rate"`
	InputPrices []float64           `json:"input_prices"`
	TaxedPrices map[string]string   `json:"taxed_prices"`
}

func (job *TaxIncluder) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, priceValue := range job.InputPrices {
		taxIncludedPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%0.2f", priceValue)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	// fmt.Println(result)

	job.TaxedPrices = result

	// filemanager.WriteJSON(fmt.Sprintf("result_%v.json", job.TaxRate*100), job)
	job.IOManager.WriteFile(job)
}

func (job *TaxIncluder) LoadData() {
	// lines, err := filemanager.ReadLines("prices.txt")
	lines, err := job.IOManager.ReadLines()
	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncluder {
	return &TaxIncluder{
		IOManager: iom,
		TaxRate:   taxRate,
		// InputPrices: []float64{10, 20, 30},
	}
}
