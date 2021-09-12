package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/symbol/{symbol}", getValueInvestBySymbol).Methods("GET")

	var port = ":8080"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func getValueInvestBySymbol(response http.ResponseWriter, request *http.Request) {
	symbol := mux.Vars(request)
	URI_INFOR := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + symbol["symbol"] + "&interval=5min&apikey=IBLB3IAD7UGS91E2"
	getUrl, err := http.Get(URI_INFOR)
	if err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(http.StatusNoContent)
	}

	responseUrl, err := ioutil.ReadAll(getUrl.Body)
	defer getUrl.Body.Close()
	if err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(http.StatusInternalServerError)
	}

	var data map[string]interface{}
	json.Unmarshal(responseUrl, &data)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
