package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	jsonMapInstance := map[string]string{
		"query": `
		{ 
			GetAllPlayers{
				id
			  }
		}
		`,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)
	if err != nil {
		fmt.Printf("There was an error marshaling the JSON instance %v", err)
	}
	// locally run graphql server
	newRequest, err := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(jsonResult))
	if err != nil {
		fmt.Printf("There was an error creating the request%v", err)
	}

	newRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(newRequest)

	if err != nil {
		fmt.Printf("There was an error executing the request%v", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Data Read Error%v", err)
	}

	fmt.Println(string(responseData))
}
