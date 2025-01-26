package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"project_sem/db"
	"project_sem/models"
	"project_sem/utils"
)

func PricesHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		connection, err := db.Connect()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		defer connection.Close()

		rows, err := db.GetAllValues(connection)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}		

		zipFile, err := utils.WriteCSVToZip(rows)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		defer os.Remove(zipFile.Name())
		defer zipFile.Close()

		response.Header().Set("Content-Type", "application/zip")
		response.Header().Set("Content-Disposition", "attachment; filename=data.zip")

		zipFile.Seek(0, 0)

		if _, err := zipFile.WriteTo(response); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if request.Method == http.MethodPost {

		err := request.ParseMultipartForm(5 << 20)
		if err != nil {
			http.Error(response, "You can upload file up to 5MB only!", http.StatusBadRequest)
			return
		}

		file, _, err := request.FormFile("file")
		if err != nil {
			http.Error(response, "Unable to read file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		buf, err := io.ReadAll(file)
		if err != nil {
			http.Error(response, "Error during file reading", http.StatusInternalServerError)
			return
		}

		rows, err := utils.GetCSVFromZip(buf)
		if err != nil {
			http.Error(response, "Unable to parse CSV file", http.StatusInternalServerError)
			return
		}

		connection, err := db.Connect()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		defer connection.Close()

		totalItems, totalCategories, totalPrice, err := db.InsertValues(connection, rows)
		if err != nil {
			http.Error(response, "Unable to save values to the database", http.StatusInternalServerError)
			return
		}

		responseValue := models.Response {
			TotalItems:		 totalItems,
			TotalCategories: totalCategories,
			TotalPrice:		 totalPrice,
		}

		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(responseValue)

	} else {
		http.Error(response, "Only GET and POST requests are allowed!", http.StatusMethodNotAllowed)
	}
}
