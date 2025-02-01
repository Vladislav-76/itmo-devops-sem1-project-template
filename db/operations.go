package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"project_sem/models"
)

func InsertValues(connection *sql.DB, rows [][]string) (int, error) {
	if len(rows) < 2 {
		return 0, fmt.Errorf("No values in CSV file")
	}

	transaction, err := connection.Begin()
	if err != nil {
		return 0, err
	}

	totalItems := 0

	for _, row := range rows[1:] {
		name := row[1]
		category := row[2]
		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}
		createDate, err := time.Parse("2006-01-02", row[4])
		if err != nil {
			transaction.Rollback()
			return 0, err
		}

		_, err = transaction.Exec(
			`INSERT INTO prices (name, category, price, create_date) VALUES ($1, $2, $3, $4)`,
			name, category, price, createDate,
		)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}

		totalItems++
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return totalItems, nil
}

func GetCategoriesAndPriceMeanings(connection *sql.DB) (int, float64, error) {
	var totalCategories int
	err := connection.QueryRow("SELECT COUNT(DISTINCT category) FROM prices").Scan(&totalCategories)
	if err != nil {
		return 0, 0, err
	}

	var totalPrice float64
	err = connection.QueryRow("SELECT SUM(price) FROM prices").Scan(&totalPrice)
	if err != nil {
		return 0, 0, err
	}
	return totalCategories, totalPrice, nil
}

func GetAllProducts(connection *sql.DB) ([]models.Product, error) {
	rows, err := connection.Query("SELECT * FROM prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.CreateDate); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
