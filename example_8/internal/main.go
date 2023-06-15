package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type Data struct {
	Status  bool   `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    Book   `json:"data" validate:"required"`
}

// Validate validates response data from external API
func (input *Data) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func main() {
	// Create an HTTP client with a timeout of 5 seconds
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	url := "http://localhost:5000/resources"
	response, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making API request: %s\n", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error parsing response JSON: %s\n", err)
		return
	}

	// Validate data from the external API
	err = data.Validate()

	if err != nil {
		fmt.Println("validation failed!")
		return
	}

	fmt.Println("Status:", data.Status)
	fmt.Println("Message:", data.Message)
	fmt.Println("Book ID:", data.Data.ID)
	fmt.Println("Book Title:", data.Data.Title)
}
