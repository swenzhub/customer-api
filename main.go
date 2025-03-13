package main

import (
	"fmt"
	"net/http"
	"os"

	"customer-api/api"
)

func main() {
	// กำหนดพอร์ต
	port := os.Getenv("PORT")
	if port == "" {
		port = "8009"
	}

	http.HandleFunc("/list", api.ListHandler)

	// เริ่ม HTTP Server
	serverAddress := ":" + port
	fmt.Println("Server is running on http://localhost" + serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		panic(err)
	}
}
