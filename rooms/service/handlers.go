package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devopsdays/gather-flag/rooms/dbclient"
	"github.com/gorilla/mux"
)

var DBClient dbclient.IMongoClient

func GetRoom(w http.ResponseWriter, r *http.Request) {

	// Read the 'roomID' path parameter from the mux map
	var roomID = mux.Vars(r)["roomID"]

	// Read the account struct BoltDB
	account, err := DBClient.QueryRoom(roomID)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
