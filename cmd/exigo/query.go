package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type searchResponse struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

var abbrevationDB map[string]string

func init() {
	abbrevationDB = make(map[string]string)
	abbrevationDB["PTO"] = "Please Turn Over"
	abbrevationDB["HP"] = "Hindustan Petroleum"
	abbrevationDB["BP"] = "Bharat Petroleum"
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// r.Queries("key", "value")
	queryString := strings.TrimSpace(r.URL.Query().Get("query"))
	fmt.Printf("Query String: %s\n", queryString)

	fullForm := abbrevationDB[queryString]
	resp := searchResponse{
		Key:   queryString,
		Value: fullForm,
	}
	json.NewEncoder(w).Encode(&resp)
}
