package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/furyGo/estuary-go-client/api"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// this file contains the following interfaces:
// ClientBuilder: Used for building executable estuary client.
// Client: A estuary client

type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
)

/**
 * Interfaces
 */

type ClientBuilder interface {
	ApiKey(apiKey string) *BuilderStruct
	Build() *ClientStruct
}

type Client interface {
	Execute(request api.Request) interface{}
}

/**
 * Structs
 */

type BuilderStruct struct {
	apiKey string
}

type ClientStruct struct {
	header map[string]string
	url    string
	method RequestMethod
	body   []byte
}

/**
 * Implementations
 */

// NewClientBuilder create a new client builder
// All starts from here
func NewClientBuilder() *BuilderStruct {
	return &BuilderStruct{}
}

func (builder *BuilderStruct) ApiKey(
	apiKey string) *BuilderStruct {
	// TODO verify apiKey
	//if !verify {
	//	panic("apiKey is not valid")
	//}
	builder.apiKey = apiKey
	return builder
}

func (builder *BuilderStruct) Build() *ClientStruct {
	headerMap := make(map[string]string)
	headerMap["Authorization"] = "Bearer " + builder.apiKey
	headerMap["Content-Type"] = "application/json"
	client := &ClientStruct{}
	client.OverrideHeaderMap(headerMap)
	return client
}

func (client *ClientStruct) OverrideHeaderMap(
	headerMap map[string]string) {
	client.header = headerMap
}

func (client *ClientStruct) AddHeader(key string, value string) {
	client.header[key] = value
}

func (client *ClientStruct) Execute(request api.Request) interface{} {
	// create request
	req, err := http.NewRequest(
		request.Method(),
		request.Url(),
		bytes.NewBuffer(request.Body()))
	if err != nil {
		log.Fatalf("client executed error: %v", err)
		return nil
	}
	// header
	for k, v := range client.header {
		req.Header.Set(k, v)
	}
	// instance net client
	netClient := &http.Client{}
	netResp, err := netClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("close body error: %v", err)
		}
	}(netResp.Body)

	// convert net response to local response
	localResp := request.ResponseInstance()
	body, _ := ioutil.ReadAll(netResp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &localResp)
	if err != nil {
		log.Fatalf("unmarshal error: %v", err)
		return nil
	}
	return localResp
}
