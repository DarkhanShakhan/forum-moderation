package http

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/config"
	postsC "github.com/DarkhanShakhan/forum-moderation/internal/ports/http/handlers/posts"
	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

func NewServer(config *config.Config, postsService posts.Service) *http.Server {
	mux := http.NewServeMux()
	m := middleware.New()
	for _, ctrl := range []Controller{
		postsC.New(postsService, m),
	} {
		ctrl.Init(mux)
	}
	// TODO: add more configs
	return &http.Server{
		Addr: config.HTTPAddress,
	}
}

type Controller interface {
	Init(mux *http.ServeMux)
}
