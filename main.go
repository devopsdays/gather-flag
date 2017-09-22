package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"html/template"
)

var appName = "web"

// Account represents the format of an account. Helpful, no?
type Account struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func showDemoAccount(w http.ResponseWriter, r *http.Request) {
	accountID := "59c3f494427d620038560f57"
	// QueryEscape escapes the account string so
	// it can be safely placed inside a URL query
	safeAccount := url.QueryEscape(accountID)

	url := fmt.Sprintf("http://accounts:8080/accounts/%s", safeAccount)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Account

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	// username := record.Username
	// firstname := record.FirstName
	// lastname := record.LastName
	// email := record.Email

	tmpl := template.Must(template.ParseFiles("accountdetail.html"))

	tmpl.Execute(w, record)

}

func main() {
	fmt.Printf("Starting %v\n", appName)
	http.HandleFunc("/", showDemoAccount)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
