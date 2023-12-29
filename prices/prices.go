package prices

import (
	"fmt"

	"github.com/horlathunbosun/price-calculator/conversion"
	"github.com/horlathunbosun/price-calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager        filemanager.FileManager
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]string
}

func (job *TaxIncludedPricesJob) loadData() {

	lines, err := job.IOManager.Readline()

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
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrice = result

	job.IOManager.WriteResult(job)

}

func NewTaxIncludedPricesJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
