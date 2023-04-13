package ai

import (
	"context"
	"errors"
	"log"

	"github.com/imroc/req/v3"
	"github.com/riba2534/let_the_bullets_fly_ai/fly/config"
	"github.com/sashabaranov/go-openai"
)

var newClient *req.Client

func Init() {
	// newClient = req.C().DevMode()
	newClient = req.C()
}

// OpenAI API 错误信息结构体
type OpenAIError struct {
	Message string `json:"message"`
	Details string `json:"details"`
	Code    int    `json:"code"`
}

func GetEmbeddings(ctx context.Context, input string) ([]float32, error) {
	var embeddingResponse openai.EmbeddingResponse
	var openAIError OpenAIError
	resp, err := newClient.R().
		SetHeader("Content-Type", "application/json").
		SetBearerAuthToken(config.C.OPENAI_API_KEY).
		SetBody(&openai.EmbeddingRequest{
			Input: []string{input},
			Model: openai.AdaEmbeddingV2,
			User:  "let_the_bullets_fly_ai",
		}).
		SetSuccessResult(&embeddingResponse).
		SetErrorResult(&openAIError).
		Post(config.C.OPENAI_URL + "/v1/embeddings")
	if err != nil {
		log.Printf("GetEmbeddings error, err=%+v, resp=%s", err, resp.Dump())
		return nil, err
	}
	if len(embeddingResponse.Data) == 0 {
		log.Printf("GetEmbeddings error, resp=%s", resp.Dump())
		return nil, errors.New("embeddings no data")
	}
	return embeddingResponse.Data[0].Embedding, nil
}
