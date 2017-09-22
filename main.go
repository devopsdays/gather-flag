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
var accountsURL = "http://accounts:8080/accounts/"
var accountService AccountService
var client *http.Client

// Account represents the format of an account. Helpful, no?
type Account struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type AccountService struct {
	Client http.Client
}

// type AccountsResult struct {
// 	Accounts []Account `json:"results"`
// 	Paging   Paging    `json:"paging"`
// }

type AccountsResult []Account

type Paging struct {
	NextPage string `json:"nextPage"`
}

// func showDemoAccount(w http.ResponseWriter, r *http.Request) {
// 	accountID := "59c3f494427d620038560f57"
// 	// QueryEscape escapes the account string so
// 	// it can be safely placed inside a URL query
// 	safeAccount := url.QueryEscape(accountID)
//
// 	url := fmt.Sprintf("http://accounts:8080/accounts/%s", safeAccount)
//
// 	// Build the request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal("NewRequest: ", err)
// 		return
// 	}
//
// 	// For control over HTTP client headers,
// 	// redirect policy, and other settings,
// 	// create a Client
// 	// A Client is an HTTP client
// 	// client := &http.Client{}
//
// 	// Send the request via a client
// 	// Do sends an HTTP request and
// 	// returns an HTTP response
// 	resp, err := accountService.getAccount(req)
// 	if err != nil {
// 		log.Fatal("Do: ", err)
// 		return
// 	}
//
// 	// Callers should close resp.Body
// 	// when done reading from it
// 	// Defer the closing of the body
// 	defer resp.Body.Close()
//
// 	// Fill the record with the data from the JSON
// 	var record Account
//
// 	// Use json.Decode for reading streams of JSON data
// 	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
// 		log.Println(err)
// 	}
//
// 	// username := record.Username
// 	// firstname := record.FirstName
// 	// lastname := record.LastName
// 	// email := record.Email
//
// 	tmpl := template.Must(template.ParseFiles("accountdetail.html"))
//
// 	tmpl.Execute(w, record)
//
// }

func accountListHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := accountService.getAccounts()

	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles("accountList.html")
	t.Execute(w, accounts)

}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := accountService.getAccount()

	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles("accountdetail.html")
	t.Execute(w, accounts)

}

func (svc AccountService) getAccounts() (result AccountsResult, err error) {

	resp, err := svc.Client.Get(accountsURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Hello")
	err = json.NewDecoder(resp.Body).Decode(&result)
	// err = json.Unmarshal([]byte(resp.Body), &result)

	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	log.Println(err)
	// }
	return
}

func (svc AccountService) getAccount() (result Account, err error) {
	accountID := "59c3f494427d620038560f57"
	// QueryEscape escapes the account string so
	// it can be safely placed inside a URL query
	safeAccount := url.QueryEscape(accountID)

	url := fmt.Sprintf("http://accounts:8080/accounts/%s", safeAccount)
	resp, err := svc.Client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	log.Println(err)
	// }
	return
}

func main() {
	// accountService = http.Client{}
	fmt.Printf("Starting %v\n", appName)
	http.HandleFunc("/", accountHandler)
	http.HandleFunc("/accounts/", accountListHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
