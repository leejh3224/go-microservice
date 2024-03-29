package service

import (
	"db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DbClient ...
var DbClient db.IBoltClient

// GetAccount ...
func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountID"]
	account, err := DbClient.QueryAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
