package categories

import (
	"context"
	"database/sql"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
)

type Repository interface {
	GetCategoriesByPostID(ctx context.Context, postID int64) ([]*entity.Category, error)
	GetPostCategoriesByPostIDs(ctx context.Context, postIDs []int64) ([]*entity.PostCategories, error)
	GetCategories(ctx context.Context) ([]*entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) (int64, error)
	DeleteCategoryByID(ctx context.Context, id int64) error
	GetCategoryByID(ctx context.Context, id int64) (*entity.Category, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

// TODO: implement
func (r *repository) GetCategoriesByPostID(ctx context.Context, postID int64) ([]*entity.Category, error) {
	return nil, nil
}

// TODO: implement
func (r *repository) GetCategories(ctx context.Context) ([]*entity.Category, error) { return nil, nil }

// TODO: implement
func (r *repository) CreateCategory(ctx context.Context, category *entity.Category) (int64, error) {
	return 0, nil
}

// TODO: implement
func (r *repository) DeleteCategoryByID(ctx context.Context, id int64) error { return nil }

// TODO: implement
func (r *repository) GetCategoryByID(ctx context.Context, id int64) (*entity.Category, error) {
	return nil, nil
}

func (r *repository) GetPostCategoriesByPostIDs(ctx context.Context, postIDs []int64) ([]*entity.PostCategories, error) {
	return nil, nil
}
