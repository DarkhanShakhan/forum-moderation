package posts

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/comments"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

const (
	getPostByIDPattern         = "/api/posts/"
	getPostByCategoryIDPattern = "/api/posts/categories/"
)

type controller struct {
	postsService    posts.Service
	commentsService comments.Service
	m               middleware.Middleware
}

func New(postsService posts.Service, commentsService comments.Service, middleware middleware.Middleware) *controller {
	return &controller{
		postsService:    postsService,
		commentsService: commentsService,
		m:               middleware,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/posts/", c.m.GET(c.m.MatchPattern(c.getPostByIDHandler, getPostByIDPattern)))
	mux.HandleFunc("/api/posts", c.m.GET(c.getPostsHandler))
	mux.HandleFunc("/api/posts/mine", c.m.GET(c.getPostsByAuthorID))
	mux.HandleFunc("/api/posts/categories/", c.m.GET(c.m.MatchPattern(c.getPostsByCategoryID, getPostByCategoryIDPattern)))
	mux.HandleFunc("/api/posts/new", c.m.POST(c.createPostHandler))
}
