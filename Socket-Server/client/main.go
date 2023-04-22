package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "http://localhost:8080/student", nil)
	if err != nil {
		fmt.Println("Error new Request... ")
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error response")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error read all...")
	}
	fmt.Println(string(body))
}
