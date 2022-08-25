package main

import (
	"estuary-go-client/api"
	"estuary-go-client/core"
	"fmt"
)

func main() {
	apiKey := ""
	client := core.NewClientBuilder().ApiKey(apiKey).Build()

	request := &api.ContentDealsRequest{}
	response := client.Execute(request)
	fmt.Println(response)
}
