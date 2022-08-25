package api

import (
	"encoding/json"
)

type Request interface {
	Url() string
	Method() string
	Body() []byte
	ResponseInstance() interface{}
}

// ContentAddIPFSRequest
// request uri https://api.estuary.tech/content/add-ipfs
// name -> file name
// root -> file cid
type ContentAddIPFSRequest struct {
	Name string `json:"name"`
	Root string `json:"root"`
}

func (request *ContentAddIPFSRequest) Url() string {
	return "https://api.estuary.tech/content/add-ipfs"
}

func (request *ContentAddIPFSRequest) Method() string {
	return "POST"
}

func (request *ContentAddIPFSRequest) Body() []byte {
	jsonStr, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	return jsonStr
}

func (request *ContentAddIPFSRequest) ResponseInstance() interface{} {
	return &ContentAddIPFSResponse{}
}

// ContentDealsRequest
// request uri https://api.estuary.tech/content/deals
type ContentDealsRequest struct {
}

func (request *ContentDealsRequest) Url() string {
	return "https://api.estuary.tech/content/deals"
}

func (request *ContentDealsRequest) Method() string {
	return "GET"
}

func (request *ContentDealsRequest) Body() []byte {
	return nil
}

func (request *ContentDealsRequest) ResponseInstance() interface{} {
	return &[]ContentDealsResponse{}
}

// ContentStatusRequest
// request uri https://api.estuary.tech/content/status/:dealId
type ContentStatusRequest struct {
	DealId string `json:"dealId"`
}

func (request *ContentStatusRequest) Url() string {
	return "https://api.estuary.tech/content/status/" + request.DealId
}

func (request *ContentStatusRequest) Method() string {
	return "GET"
}

func (request *ContentStatusRequest) Body() []byte {
	return nil
}

func (request *ContentStatusRequest) ResponseInstance() interface{} {
	return &ContentStatusResponse{}
}
