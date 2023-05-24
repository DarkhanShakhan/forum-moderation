package categories

import (
	"context"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/categories"
)

type Service interface {
	CreateCategory(ctx context.Context, categoryTitle string) (int64, error)
	DeleteCategory(ctx context.Context, categoryID int64) error
	GetCategories(ctx context.Context) ([]*entity.Category, error)
	GetCategory(ctx context.Context, categoryID int64) (*entity.Category, error)
}

type service struct {
	categoriesRepository categories.Repository
}

func New(categoriesRepository categories.Repository) Service {
	return &service{
		categoriesRepository: categoriesRepository,
	}
}

func (s *service) CreateCategory(ctx context.Context, categoryTitle string) (int64, error) {
	return s.categoriesRepository.CreateCategory(ctx, &entity.Category{
		Title: categoryTitle,
	})
}

func (s *service) DeleteCategory(ctx context.Context, categoryID int64) error {
	return s.categoriesRepository.DeleteCategoryByID(ctx, categoryID)
}

func (s *service) GetCategories(ctx context.Context) ([]*entity.Category, error) {
	return s.categoriesRepository.GetCategories(ctx)
}

func (s *service) GetCategory(ctx context.Context, categoryID int64) (*entity.Category, error) {
	return s.categoriesRepository.GetCategoryByID(ctx, categoryID)
}
