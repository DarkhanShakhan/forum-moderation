package admin

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/comments"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

const (
	deleteCategoryPattern = "/api/admin/categories/"
	deleteCommentPattern  = "/apo/admin/comments/"
)

type controller struct {
	postsService      posts.Service
	categoriesService categories.Service
	commentsService   comments.Service
	m                 middleware.Middleware
}

func New(postsService posts.Service, categoriesService categories.Service, commentsService comments.Service, middleware middleware.Middleware) *controller {
	return &controller{
		postsService:      postsService,
		categoriesService: categoriesService,
		commentsService:   commentsService,
		m:                 middleware,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	// TODO: add admin middleware
	mux.HandleFunc("/api/admin/categories/new", c.m.POST(c.createCategoryHandler))
	mux.HandleFunc("/api/admin/categories/", c.m.DELETE(
		c.m.MatchPattern(c.deleteCategoryHandler, deleteCategoryPattern)))
	mux.HandleFunc("/api/admin/comments/", c.m.DELETE(
		c.m.MatchPattern(c.deleteCommentHandler, deleteCommentPattern),
	))
	mux.HandleFunc("api/admin/comments/set_visible", c.m.POST(c.setCommentVisibleHandler))
}
