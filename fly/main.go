package main

import (
	"io"
	"log"
	"os"

	"github.com/riba2534/let_the_bullets_fly_ai/fly/config"
	"github.com/riba2534/let_the_bullets_fly_ai/fly/dal/ai"
	"github.com/riba2534/let_the_bullets_fly_ai/fly/dal/qdrant"
)

func init() {
	// 初始化log
	f, _ := os.OpenFile("run.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.SetPrefix("[let_the_bullets_fly_ai] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 初始化配置
	config.Init()
	// 初始化 dal
	ai.Init()
	qdrant.Init()

}

func main() {

}
