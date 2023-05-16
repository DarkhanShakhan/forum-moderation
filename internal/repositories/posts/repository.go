package posts

import (
	"context"
	"database/sql"
	"time"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	"github.com/DarkhanShakhan/forum-moderation/internal/errors"
)

type Repository interface {
	GetPostByID(ctx context.Context, id int64) (*entity.Post, error)
	GetPosts(ctx context.Context) ([]*entity.Post, error)
	GetPostsByCategory(ctx context.Context, categoryID int64) ([]*entity.Post, error)
	GetPostsByAuthorID(ctx context.Context, authorID int64) ([]*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) (int64, error)
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage string) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetPostByID(ctx context.Context, id int64) (*entity.Post, error) {
	var m model
	rows, err := r.db.QueryContext(ctx, getPostByIDStmt, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&m.ID, &m.Title, &m.Content, &m.AuthorID, &m.LikesCount, &m.DislikesCount)
		return m.convert(), nil
	}
	return nil, errors.ErrPostNotFound
}

func (r *repository) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	var (
		mm models
		m  model
	)
	rows, err := r.db.QueryContext(ctx, getPostsStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&m.ID, &m.Title, &m.Content, &m.AuthorID, &m.LikesCount, &m.DislikesCount)
		mm = append(mm, &m)
	}

	return mm.convert(), nil
}

func (r *repository) GetPostsByCategory(ctx context.Context, categoryID int64) ([]*entity.Post, error) {
	var (
		mm models
		m  model
	)
	rows, err := r.db.QueryContext(ctx, getPostsByCategoryStmt, categoryID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&m.ID, &m.Title, &m.Content, &m.AuthorID, &m.LikesCount, &m.DislikesCount)
		mm = append(mm, &m)
	}

	return mm.convert(), nil
}

func (r *repository) GetPostsByAuthorID(ctx context.Context, authorID int64) ([]*entity.Post, error) {
	var (
		mm models
		m  model
	)
	rows, err := r.db.QueryContext(ctx, getPostsByAuthorIDStmt, authorID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&m.ID, &m.Title, &m.Content, &m.AuthorID, &m.LikesCount, &m.DislikesCount)
		mm = append(mm, &m)
	}

	return mm.convert(), nil
}

func (r *repository) CreatePost(ctx context.Context, post *entity.Post) (int64, error) {
	res, err := r.db.ExecContext(ctx, createPostStmt, post.Title, post.Content, post.AuthorID, time.Now())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *repository) UpdatePost(ctx context.Context, post *entity.Post) error {
	return nil
}

func (r *repository) DeletePost(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage string) error {
	res, err := r.db.ExecContext(ctx, deletePostStmt, time.Now(), deleteMessage, deleteCategory, id)
	if err != nil {
		return err
	}
	rAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rAffected == 0 {
		return errors.ErrPostNotFound
	}
	return nil
}

func (r *repository) SetVisible(ctx context.Context, id uint64, visible bool) error {
	return nil
}

func (r *repository) ReportPost(ctx context.Context, id uint64, reportCategory enum.ReportCategory, reportMessage string) error {
	return nil
}

func (r *repository) GetReportedPosts(ctx context.Context) ([]*entity.Post, error) {
	return nil, nil
}

func (r *repository) GetDeletedPosts(ctx context.Context) ([]*entity.Post, error) {
	return nil, nil
}
