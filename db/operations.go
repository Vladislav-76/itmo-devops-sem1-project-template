package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func InsertValues(connection *sql.DB, rows [][]string) (int, int, int, error) {
	if len(rows) < 2 {
		return 0, 0, 0, fmt.Errorf("No values in CSV file")
	}

	transaction, err := connection.Begin()
	if err != nil {
		return 0, 0, 0, err
	}
	defer transaction.Rollback()

	totalItems := 0
	categoryMap := make(map[string]struct{})
	totalPrice := 0.0

	for _, row := range rows[1:] {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return 0, 0, 0, err
		}
		name := row[1]
		category := row[2]
		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return 0, 0, 0, err
		}
		createDate, err := time.Parse("2006-01-02", row[4])
		if err != nil {
			return 0, 0, 0, err
		}

		_, err = transaction.Exec(
			`INSERT INTO prices (id, name, category, price, create_date) VALUES ($1, $2, $3, $4, $5)`,
			id, name, category, price, createDate,
		)
		if err != nil {
			return 0, 0, 0, err
		}

		totalItems++
		categoryMap[category] = struct{}{}
		totalPrice += price
	}

	err = transaction.Commit()
	if err != nil {
		return 0, 0, 0, err
	}

	return totalItems, len(categoryMap), int(totalPrice), nil
}

func GetAllValues(connection *sql.DB) (*sql.Rows, error) {
	rows, err := connection.Query("SELECT * FROM prices")
	if err != nil {
		return nil, err
	}
	
	return rows, nil
}
