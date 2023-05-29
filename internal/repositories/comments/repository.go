package comments

import (
	"context"
	"database/sql"
	"time"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	"github.com/DarkhanShakhan/forum-moderation/internal/errors"
)

type Repository interface {
	CreateComment(ctx context.Context, comment *entity.Comment) (int64, error)
	GetCommentsByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error)
	DeleteCommentByID(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage *string) error
	SetVisible(ctx context.Context, id int64, visible bool) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateComment(ctx context.Context, comment *entity.Comment) (int64, error) {
	res, err := r.db.ExecContext(ctx, createCommentStmt, comment.PostID, comment.AuthorID, comment.Content, time.Now())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *repository) GetCommentsByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error) {
	var (
		mm models
		m  model
	)
	rows, err := r.db.QueryContext(ctx, getCommentsByPostIDStmt, postID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&m.ID, &m.PostID, &m.Content, &m.AuthorID, &m.LikesCount, &m.DislikesCount)
		mm = append(mm, &m)
	}

	return mm.convert(), nil
}

func (r *repository) DeleteCommentByID(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage *string) error {
	res, err := r.db.ExecContext(ctx, deleteCommentStmt, time.Now(), deleteMessage, deleteCategory, id)
	if err != nil {
		return err
	}
	rAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rAffected == 0 {
		return errors.ErrCommentNotFound
	}
	return nil
}

func (r *repository) SetVisible(ctx context.Context, id int64, visible bool) error {
	res, err := r.db.ExecContext(ctx, setVisibleStmt, visible, id)
	if err != nil {
		return err
	}
	rAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rAffected == 0 {
		return errors.ErrCommentNotFound
	}
	return nil
}
