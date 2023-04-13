package qdrant

import (
	"log"

	"github.com/imroc/req/v3"
	"github.com/riba2534/let_the_bullets_fly_ai/fly/config"
)

var newClient *req.Client

func Init() {
	newClient = req.C().DevMode()
}

type CreateCollectioRequest struct {
	Name    string  `json:"name"`
	Vectors Vectors `json:"vectors"`
}

type Vectors struct {
	Size     int    `json:"size"`
	Distance string `json:"distance"`
}

type CreateCollectionResponse struct {
	Result bool    `json:"result"`
	Status string  `json:"status"`
	Time   float64 `json:"time"`
}

// 创建集合
func CreateCollection(collectionName string) bool {
	var resp CreateCollectionResponse
	r, err := newClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("api-key", config.C.QDRANT_API_KEY).
		SetPathParam("collection_name", collectionName).
		SetBody(&CreateCollectioRequest{
			Name: collectionName,
			Vectors: Vectors{
				Size:     1536,
				Distance: "Cosine",
			},
		}).
		SetSuccessResult(&resp).
		Put(config.C.QDRANT_URL + config.C.QDRANT_PORT + "/collections/{collection_name}")
	if err != nil {
		log.Printf("collections error, err=%+v, resp=%s", err, r.Dump())
		return false
	}
	return resp.Result
}

// 删除集合
func DeleteCollection(collectionName string) bool {
	var resp CreateCollectionResponse
	r, err := newClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("api-key", config.C.QDRANT_API_KEY).
		SetPathParam("collection_name", collectionName).
		SetSuccessResult(&resp).
		Delete(config.C.QDRANT_URL + config.C.QDRANT_PORT + "/collections/{collection_name}")
	if err != nil {
		log.Printf("collections error, err=%+v, resp=%s", err, r.Dump())
		return false
	}
	return resp.Result
}
