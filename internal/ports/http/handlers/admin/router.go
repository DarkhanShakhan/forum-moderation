package admin

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

type controller struct {
	postsService      posts.Service
	categoriesService categories.Service
	m                 middleware.Middleware
}

func New(postsService posts.Service, categoriesService categories.Service, middleware middleware.Middleware) *controller {
	return &controller{
		postsService:      postsService,
		categoriesService: categoriesService,
		m:                 middleware,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	// TODO: add admin middleware
	mux.HandleFunc("/api/admin/categories/new", c.m.POST(c.createCategoryHandler))
}
