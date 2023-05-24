package posts

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

type controller struct {
	postsService posts.Service
	m            middleware.Middleware
}

func New(postsService posts.Service, middleware middleware.Middleware) *controller {
	return &controller{
		postsService: postsService,
		m:            middleware,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/posts/", c.m.GET(c.getPostByIDHandler))
	mux.HandleFunc("/api/posts", c.m.GET(c.getPostsHandler))
	mux.HandleFunc("/api/posts/mine", c.m.GET(c.getPostsByAuthorID))
	mux.HandleFunc("/api/posts/categories/", c.m.GET(c.getPostsByCategoryID))
	mux.HandleFunc("/api/posts/new", c.m.POST(c.createPostHandler))
}
