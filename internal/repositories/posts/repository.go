package posts

import (
	"context"
	"database/sql"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	"github.com/DarkhanShakhan/forum-moderation/internal/errors"
)

type Repository interface {
	GetPostByID(ctx context.Context, id uint64) (*entity.Post, error)
	GetPosts(ctx context.Context) ([]*entity.Post, error)
	GetPostsByCategory(ctx context.Context, categoryID uint64) ([]*entity.Post, error)
	GetPostsByAuthorID(ctx context.Context, authorID uint64) ([]*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) (uint64, error)
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id uint64) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetPostByID(ctx context.Context, id uint64) (*entity.Post, error) {
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
		mm = append(mm, m)
	}

	return mm.convert(), nil
}

func (r *repository) GetPostsByCategory(ctx context.Context, categoryID uint64) ([]*entity.Post, error) {
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
		mm = append(mm, m)
	}

	return mm.convert(), nil
}

func (r *repository) GetPostsByAuthorID(ctx context.Context, authorID uint64) ([]*entity.Post, error) {
	return nil, nil
}

func (r *repository) CreatePost(ctx context.Context, post *entity.Post) (uint64, error) {
	return 0, nil
}

func (r *repository) UpdatePost(ctx context.Context, post *entity.Post) error {
	return nil
}

func (r *repository) DeletePost(ctx context.Context, id uint64, deleteCategory enum.ReportCategory, deleteMessage string) error {
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
