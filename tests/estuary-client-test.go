package main

import (
	"fmt"
	"github.com/furyGo/estuary-go-client/api"
	"github.com/furyGo/estuary-go-client/core"
)

func main() {
	apiKey := ""
	client := core.NewClientBuilder().ApiKey(apiKey).Build()

	request := &api.ContentDealsRequest{}
	response := client.Execute(request)
	fmt.Println(response)
}
