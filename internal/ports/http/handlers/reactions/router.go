package reactions

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/ports/http/middleware"
	"github.com/DarkhanShakhan/forum-moderation/internal/services/reactions"
)

type controller struct {
	reactionsService reactions.Service
	m                middleware.Middleware
}

func New(reactionsService reactions.Service, m middleware.Middleware) *controller {
	return &controller{
		reactionsService: reactionsService,
		m:                m,
	}
}

func (c *controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/posts/reactions", c.m.POST(c.upsertPostReactions))
	mux.HandleFunc("/api/comments/reactions", c.m.POST(c.upsertCommentReactions))
}
