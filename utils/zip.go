package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
)

func GetCSVFromZip(file []byte) ([][]string, error) {
	reader, err := zip.NewReader(bytes.NewReader(file), int64(len(file)))
	if err != nil {
		return nil, err
	}

	for _, file := range reader.File {
		if strings.HasSuffix(file.Name, ".csv") {
			csvFile, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer csvFile.Close()

			return ReadCSV(csvFile)
		}
	}
	return nil, fmt.Errorf("No CSV file in zip file!")
}