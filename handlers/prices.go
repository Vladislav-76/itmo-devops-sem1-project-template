package handlers

import (
	"io"
	"net/http"
)

func PricesHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		response.WriteHeader(http.StatusOK)
		response.Write([]byte("GET apiPrices"))
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

		response.Write([]byte("POST apiPrices"))
	} else {
		http.Error(response, "Only GET and POST requests are allowed!", http.StatusMethodNotAllowed)
	}
}
