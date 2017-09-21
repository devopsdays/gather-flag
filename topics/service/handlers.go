package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devopsdays/gather-flag/topics/dbclient"
	"github.com/gorilla/mux"
)

var DBClient dbclient.IMongoClient

func GetTopic(w http.ResponseWriter, r *http.Request) {

	// Read the 'topicID' path parameter from the mux map
	var topicID = mux.Vars(r)["topicID"]

	// Read the topic struct BoltDB
	topic, err := DBClient.QueryTopic(topicID)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(topic)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
