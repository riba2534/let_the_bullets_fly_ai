package qdrant

import (
	"testing"

	"github.com/riba2534/let_the_bullets_fly_ai/fly/config"
)

func TestCreateCollection(t *testing.T) {
	Init()
	config.Init()
	name := "test_1"
	t.Log(CreateCollection(name))
	t.Log(DeleteCollection(name))
}
