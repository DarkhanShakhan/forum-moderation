package comments

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/comments"
)

type controller struct {
	commentsService comments.Service
	m               middleware.Middleware
}

func New(commentsService comments.Service, m middleware.Middleware) *controller {
	return &controller{
		commentsService: commentsService,
		m:               m,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("api/comments/new", c.m.POST(c.createCommentHandler))
}
