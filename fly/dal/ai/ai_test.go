package ai

import (
	"context"
	"testing"

	"github.com/riba2534/let_the_bullets_fly_ai/fly/config"
)

func TestGetEmbeddings(t *testing.T) {
	Init()
	config.Init()
	ctx := context.Background()
	input := "The food was delicious and the waiter..."
	e, err := GetEmbeddings(ctx, input)
	if err != nil {
		t.Errorf("GetEmbeddings err, err=%+v", err)
		return
	}
	t.Log(e[:10])
}
