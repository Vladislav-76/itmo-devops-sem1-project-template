package utils

import (
    "encoding/csv"
    "io"
)

func ReadCSV(csvFile io.Reader) ([][]string, error) {
    csvReader := csv.NewReader(csvFile)
    return csvReader.ReadAll()
}
