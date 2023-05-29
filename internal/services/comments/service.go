package comments

import (
	"context"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/comments"
)

type Service interface {
	GetCommentsByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error)
	CreateComment(ctx context.Context, comment *entity.Comment) (int64, error)
	DeleteComment(ctx context.Context, id int64, deleteCategory string, deleteMessage *string) error
	SetVisible(ctx context.Context, id int64, visible bool) error
}

type service struct {
	commentsRepository comments.Repository
}

func New(commentsRepository comments.Repository) Service {
	return &service{
		commentsRepository: commentsRepository,
	}
}

func (s *service) GetCommentsByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error) {
	return s.commentsRepository.GetCommentsByPostID(ctx, postID)
}

func (s *service) CreateComment(ctx context.Context, comment *entity.Comment) (int64, error) {
	return s.commentsRepository.CreateComment(ctx, comment)
}

func (s *service) DeleteComment(ctx context.Context, id int64, deleteCategory string, deleteMessage *string) error {
	delCat, _ := enum.ParseStringToReportCategory(deleteCategory)
	return s.commentsRepository.DeleteCommentByID(ctx, id, delCat, deleteMessage)
}

func (s *service) SetVisible(ctx context.Context, id int64, visible bool) error {
	return s.commentsRepository.SetVisible(ctx, id, visible)
}
