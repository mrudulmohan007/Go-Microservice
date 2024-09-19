package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Test User Service
	resp, err := http.Get("http://localhost:8081/users/1")
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("User Service is working: Status Code 200")
	} else {
		fmt.Printf("User Service returned status code: %d\n", resp.StatusCode)
	}

	// Test Product Service
	resp, err = http.Get("http://localhost:8082/products/1")
	if err != nil {
		fmt.Printf("Failed to get product: %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Product Service is working: Status Code 200")
	} else {
		fmt.Printf("Product Service returned status code: %d\n", resp.StatusCode)
	}
}
