package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"project_sem/models"
)

func ReadCSV(csvFile io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

func WriteCSVToZip(products []models.Product) (*os.File, error) {
	csvFile, err := os.CreateTemp("", "data-*.csv")
	if err != nil {
		return nil, err
	}
	defer os.Remove(csvFile.Name())
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	err = writer.Write([]string{"id", "name", "category", "price", "create_date"})
	if err != nil {
		return nil, err
	}

	for _, product := range products  {

		row := []string{
			fmt.Sprintf("%d", product.ID),
			product.Name,
			product.Category,
			fmt.Sprintf("%.2f", product.Price),
			product.CreateDate.Format("2006-01-02"),
		}

		if err := writer.Write(row); err != nil {
			return nil, err
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			return nil, err
		}
	}

	zipFile, err := ZipCSV(csvFile)
	if err != nil {
		return nil, err
	}

	return zipFile, nil
}
