package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Customers []Customer `json:"customerList"`
	Total     int        `json:"totalCustomer"`
}

func List(branch string) Response {
	if strings.ToLower(branch) == "kbtg" {
		customers := []Customer{
			{Name: "John", Age: 25},
			{Name: "Jane", Age: 40},
		}
		return Response{Customers: customers, Total: len(customers)}
	}
	return Response{Customers: []Customer{}, Total: 0}

}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	branch := r.URL.Query().Get("branch")

	response := List(branch)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
