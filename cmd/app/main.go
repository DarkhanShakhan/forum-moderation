package main

import (
	"log"

	"github.com/DarkhanShakhan/forum-moderation/internal/app"
)

func main() {
	a := app.InitApp()
	defer func() {
		if err := a.CloseDB(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := a.Start(); err != nil {
		log.Fatal("app start", err)
	}
}
