package posts

import "github.com/DarkhanShakhan/forum-moderation/internal/repositories/posts"

type Service struct {
	postsRepository posts.Repository
}

func New(postsRepository posts.Repository) *Service {
	return &Service{
		postsRepository: postsRepository,
	}
}
