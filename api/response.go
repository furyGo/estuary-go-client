package api

import "time"

type Response interface {
	GetError() string
	GetStatusCode() int
	GetContent() Response
}

// ContentAddIPFSResponse
//{
// "requestid": "35478736",
// "status": "queued",
// "created": "2022-08-18T09:47:14.279767505Z",
// "delegates": [
//  "/ip4/3.134.223.177/tcp/6745/p2p/12D3KooWN8vAoGd6eurUSidcpLYguQiGZwt4eVgDvbgaS7kiGTup"
// ],
// "info": {},
// "pin": {
//  "cid": "bafybeidj7c2e3daplalccukbps4eze7473gyshspev76xi4sjfmfkuaofe",
//  "name": "bafybeidj7c2e3daplalccukbps4eze7473gyshspev76xi4sjfmfkuaofe",
//  "origins": [],
//  "meta": {}
// }
//}
type ContentAddIPFSResponse struct {
	RequestId string                 `json:"requestid"`
	Status    string                 `json:"status"`
	Created   string                 `json:"created"`
	Delegates []string               `json:"delegates"`
	Info      map[string]interface{} `json:"info"`
	Pin       ContentAddIPFSPin      `json:"pin"`
}

type ContentAddIPFSPin struct {
	Cid     string                 `json:"cid"`
	Name    string                 `json:"name"`
	Origins []interface{}          `json:"origins"`
	Meta    map[string]interface{} `json:"meta"`
}

func (response *ContentAddIPFSResponse) GetError() string {
	return ""
}

func (response *ContentAddIPFSResponse) GetStatusCode() int {
	return 200
}

func (response *ContentAddIPFSResponse) GetContent() Response {
	return response
}

// ContentDealsResponse
// {
// "0": {
//  "id": 33421266,
//  "updatedAt": "2022-07-25T05:03:08.327134Z",
//  "cid": "QmWj3UnfuDKuygzBYt2sL8xCNnBsY6WDiSsvNaGeaoiFYw",
//  "name": "aggregate",
//  "userId": 426,
//  "description": "",
//  "size": 4325553465,
//  "type": 0,
//  "active": true,
//  "offloaded": false,
//  "replication": 6,
//  "aggregatedIn": 0,
//  "aggregate": true,
//  "pinning": false,
//  "pinMeta": "",
//  "replace": false,
//  "origins": "",
//  "failed": false,
//  "location": "SHUTTLE1d45aa63-0927-451f-85d8-748d0a8e1c39HANDLE",
//  "dagSplit": false,
//  "splitFrom": 0,
//  "aggregatedFiles": 34
// }
//}
type ContentDealsResponse struct {
	Id              int64  `json:"id"`
	UpdatedAt       string `json:"updatedAt"`
	Cid             string `json:"cid"`
	Name            string `json:"name"`
	UserId          int64  `json:"userId"`
	Description     string `json:"description"`
	Size            int64  `json:"size"`
	Type            int8   `json:"type"`
	Active          bool   `json:"active"`
	Offloaded       bool   `json:"offloaded"`
	Replication     int8   `json:"replication"`
	AggregatedIn    int64  `json:"aggregatedIn"`
	Aggregate       bool   `json:"aggregate"`
	Pinning         bool   `json:"pinning"`
	PinMeta         string `json:"pinMeta"`
	Replace         bool   `json:"replace"`
	Origins         string `json:"origins"`
	Failed          bool   `json:"failed"`
	Location        string `json:"location"`
	DagSplit        bool   `json:"dagSplit"`
	SplitFrom       int64  `json:"splitFrom"`
	AggregatedFiles int64  `json:"aggregatedFiles"`
}

func (response *ContentDealsResponse) GetError() string {
	return ""
}

func (response *ContentDealsResponse) GetStatusCode() int {
	return 200
}

func (response *ContentDealsResponse) GetContent() Response {
	return response
}

// ContentStatusResponse
// {
// "content": {
//  "id": 35406719,
//  "updatedAt": "2022-08-17T05:14:19.687977Z",
//  "cid": "bafybeiagfn7g3ibn4jc72nccf7cq4rs5k64thm7aurjgkort6jaqdh2qey",
//  "name": "JetBrains.Rider-2022.2.1.dmg",
//  "userId": 426,
//  "description": "",
//  "size": 954812218,
//  "type": 0,
//  "active": true,
//  "offloaded": false,
//  "replication": 6,
//  "aggregatedIn": 33421266,
//  "aggregate": false,
//  "pinning": false,
//  "pinMeta": "",
//  "replace": false,
//  "origins": "",
//  "failed": false,
//  "location": "SHUTTLE1d45aa63-0927-451f-85d8-748d0a8e1c39HANDLE",
//  "dagSplit": false,
//  "splitFrom": 0
// },
// "deals": [],
// "failuresCount": 0
//}
type ContentStatusResponse struct {
	Content       ContentStatusContentResponse `json:"content"`
	Deals         []ContentStatusDeals         `json:"deals"`
	FailuresCount int64                        `json:"failuresCount"`
}

type ContentStatusDeals struct {
	CreatedAt           time.Time `json:"CreatedAt"`
	DeletedAt           time.Time `json:"DeletedAt"`
	Id                  int64     `json:"ID"`
	UpdatedAt           time.Time `json:"UpdatedAt"`
	Content             int64     `json:"content"`
	DealId              int64     `json:"dealId"`
	DealUuid            string    `json:"dealUuid"`
	DealProtocolVersion string    `json:"deal_protocol_version"`
	DtChan              string    `json:"dtChan"`
	Failed              bool      `json:"failed"`
	FailedAt            time.Time `json:"failedAt"`
	Miner               string    `json:"miner"`
	MinerVersion        string    `json:"miner_version"`
	OnChainAt           time.Time `json:"onChainAt"`
	PropCid             string    `json:"propCid"`
	SealedAt            time.Time `json:"sealedAt"`
	Slashed             bool      `json:"slashed"`
	TransferFinished    time.Time `json:"transferFinished"`
	TransferStarted     time.Time `json:"transferStarted"`
	UserId              int64     `json:"user_id"`
	Verified            bool      `json:"verified"`
}

type ContentStatusContentResponse struct {
	Id           int64  `json:"id"`
	UpdatedAt    string `json:"updatedAt"`
	Cid          string `json:"cid"`
	Name         string `json:"name"`
	UserId       int64  `json:"userId"`
	Description  string `json:"description"`
	Size         int64  `json:"size"`
	Type         int8   `json:"type"`
	Active       bool   `json:"active"`
	Offloaded    bool   `json:"offloaded"`
	Replication  int8   `json:"replication"`
	AggregatedIn int64  `json:"aggregatedIn"`
	Aggregate    bool   `json:"aggregate"`
	Pinning      bool   `json:"pinning"`
	PinMeta      string `json:"pinMeta"`
	Replace      bool   `json:"replace"`
	Origins      string `json:"origins"`
	Failed       bool   `json:"failed"`
	Location     string `json:"location"`
	DagSplit     bool   `json:"dagSplit"`
	SplitFrom    int64  `json:"splitFrom"`
}

func (response *ContentStatusResponse) GetError() string {
	return ""
}

func (response *ContentStatusResponse) GetStatusCode() int {
	return 200
}

func (response *ContentStatusResponse) GetContent() Response {
	return response
}
