package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

func ReadCSV(csvFile io.Reader) ([][]string, error) {
    csvReader := csv.NewReader(csvFile)
    return csvReader.ReadAll()
}

func WriteCSVToZip(rows *sql.Rows) (*os.File, error) {
    csvFile, err := os.CreateTemp("", "data-*.csv")
    if err != nil {
        return nil, err
    }
    defer rows.Close() 
    defer os.Remove(csvFile.Name())
    defer csvFile.Close()

    writer := csv.NewWriter(csvFile)
    err = writer.Write([]string{"id", "name", "category", "price", "create_date"})
    if err != nil {
        return nil, err
    }

    for rows.Next() {
		var id int
		var name, category string
		var price float64
		var createDate time.Time

        if err := rows.Scan(&id, &name, &category, &price, &createDate); err != nil {
            return nil, err
        }

        row := []string{
			fmt.Sprintf("%d", id),
			name,
			category,
			fmt.Sprintf("%.2f", price),
			createDate.Format("2006-01-02"),
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