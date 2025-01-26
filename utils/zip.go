package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
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

func ZipCSV(csvFile *os.File) (*os.File, error) {
	zipFile, err := os.CreateTemp("", "data-*.zip")
	if err != nil {
		return nil, err
	}
	
	zipWriter := zip.NewWriter(zipFile)
	fileInZip, err := zipWriter.Create("data.csv")
	if err != nil {
		return nil, err
	}

	csvFile.Seek(0, 0)
	if _, err := csvFile.WriteTo(fileInZip); err != nil {
		return nil, err
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return zipFile, nil
}