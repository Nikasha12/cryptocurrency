// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var apiURL = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=cad"

func getPrice(cryptoName string) (float64, error) {
	url := fmt.Sprintf(apiURL, cryptoName)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}

	var data map[string]map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return 0, err
	}

	priceData, ok := data[cryptoName]
	if !ok {
		return 0, fmt.Errorf("data for %s not found in response", cryptoName)
	}

	cadPrice, ok := priceData["cad"]
	if !ok {
		return 0, fmt.Errorf("CAD price not found for %s", cryptoName)
	}

	// Assuming the CAD price is always a float64, you may need additional checks
	cadPriceFloat, ok := cadPrice.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected data type for CAD price for %s", cryptoName)
	}

	return cadPriceFloat, nil
}

func priceHandler(w http.ResponseWriter, r *http.Request) {
	cryptoName := r.URL.Query().Get("crypto")
	if cryptoName == "" {
		http.Error(w, "crypto parameter is required", http.StatusBadRequest)
		return
	}

	price, err := getPrice(cryptoName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Current price of %s: %.2f CAD", cryptoName, price)
}

func main() {
	http.HandleFunc("/price", priceHandler)
	fmt.Println("Server is listening on port 9090")
	http.ListenAndServe(":9090", nil)
}
