package prices

import (
	"fmt"

	"github.com/horlathunbosun/price-calculator/conversion"
	"github.com/horlathunbosun/price-calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPricesJob) loadData() error {

	lines, err := job.IOManager.Readline()

	if err != nil {
		return err
	}
	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil

}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.loadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrice = result

	return job.IOManager.WriteResult(job)

}

func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
