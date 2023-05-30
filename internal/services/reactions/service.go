package reactions

import (
	"context"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/reactions"
)

type Service interface {
	UpsertPostReaction(ctx context.Context, postReaciton *entity.Reaction) error
	UpsertCommentReaction(ctx context.Context, commentReaction *entity.Reaction) error
}

type service struct {
	reactionsRepository reactions.Repository
}

func New(reactionsRepository reactions.Repository) Service {
	return &service{
		reactionsRepository: reactionsRepository,
	}
}

func (s *service) UpsertPostReaction(ctx context.Context, postReaciton *entity.Reaction) error {
	return s.reactionsRepository.UpsertPostReaction(ctx, postReaciton)
}

func (s *service) UpsertCommentReaction(ctx context.Context, commentReaction *entity.Reaction) error {
	return s.reactionsRepository.UpsertCommentReaction(ctx, commentReaction)
}
