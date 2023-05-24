package http

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/config"
	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/handlers/admin"
	postsC "github.com/DarkhanShakhan/forum-moderation/internal/ports/http/handlers/posts"
	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

func NewServer(config *config.Config, postsService posts.Service, categoriesService categories.Service) *http.Server {
	mux := http.NewServeMux()
	m := middleware.New()
	for _, ctrl := range []Controller{
		postsC.New(postsService, m),
		admin.New(postsService, categoriesService, m),
	} {
		ctrl.Init(mux)
	}
	// TODO: add more configs
	return &http.Server{
		Addr:    config.HTTPAddress,
		Handler: mux,
	}
}

type Controller interface {
	Init(mux *http.ServeMux)
}
