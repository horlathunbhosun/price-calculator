package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/horlathunbosun/price-calculator/conversion"
)

type TaxIncludedPricesJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludedPricesJob) loadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Unable to open file")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Convert price to float filed")
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()

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
