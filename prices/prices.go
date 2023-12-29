package prices

import (
	"fmt"

	"github.com/horlathunbosun/price-calculator/conversion"
	"github.com/horlathunbosun/price-calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludedPricesJob) loadData() {

	lines, err := filemanager.Readline("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Convert price to float filed")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPricesJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.0f", price)] = fmt.Sprintf("%.0f", taxIncludedPrice)
	}
	fmt.Println(result)

}

func NewTaxIncludedPricesJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
