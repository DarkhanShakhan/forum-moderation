package main

import (
	"log"

	"github.com/DarkhanShakhan/forum-moderation/internal/app"
)

func main() {
	a := app.InitApp()
	if err := a.Start(); err != nil {
		log.Fatal("app start", err)
	}
}
