package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Convert price to float filed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
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
