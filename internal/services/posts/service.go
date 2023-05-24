package posts

import (
	"context"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	"github.com/DarkhanShakhan/forum-moderation/internal/errors"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/posts"
	errgroup "github.com/DarkhanShakhan/forum-moderation/internal/util"
)

type Service interface {
	GetPostByID(ctx context.Context, id int64) (*entity.Post, error)
	GetPosts(ctx context.Context) ([]*entity.Post, error)
	GetPostsByAuthorID(ctx context.Context, authorID int64) ([]*entity.Post, error)
	GetPostsByCategory(ctx context.Context, categoryID int64) (*entity.CategoryPosts, error)
	CreatePost(ctx context.Context, post *entity.Post) (int64, error)
	DeletePost(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage string) error
}

type service struct {
	postsRepository      posts.Repository
	categoriesRepository categories.Repository
}

func New(postsRepository posts.Repository, categoriesRepository categories.Repository) Service {
	return &service{
		postsRepository:      postsRepository,
		categoriesRepository: categoriesRepository,
	}
}

func (s *service) GetPostByID(ctx context.Context, id int64) (*entity.Post, error) {
	var (
		post *entity.Post
		cats []*entity.Category
		err  error
	)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(
		func() error {
			post, err = s.postsRepository.GetPostByID(ctx, id)
			return err
		})
	g.Go(
		func() error {
			cats, err = s.categoriesRepository.GetCategoriesByPostID(ctx, id)
			return err
		})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	post.Categories = cats
	return post, nil
}

func (s *service) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	posts, err := s.postsRepository.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return s.mergeCategories(ctx, posts)
}

func (s *service) GetPostsByAuthorID(ctx context.Context, authorID int64) ([]*entity.Post, error) {
	posts, err := s.postsRepository.GetPostsByAuthorID(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return s.mergeCategories(ctx, posts)
}

func (s *service) GetPostsByCategory(ctx context.Context, categoryID int64) (*entity.CategoryPosts, error) {
	category, err := s.categoriesRepository.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, errors.ErrCategoryNotFound
	}
	posts, err := s.postsRepository.GetPostsByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	posts, err = s.mergeCategories(ctx, posts)
	if err != nil {
		return nil, err
	}
	return &entity.CategoryPosts{
		Category: category,
		Posts:    posts,
	}, nil
}

func (s *service) CreatePost(ctx context.Context, post *entity.Post) (int64, error) {
	return s.postsRepository.CreatePost(ctx, post)
}

func (s *service) DeletePost(ctx context.Context, id int64, deleteCategory enum.ReportCategory, deleteMessage string) error {
	return s.postsRepository.DeletePost(ctx, id, deleteCategory, deleteMessage)
}

func (s *service) mergeCategories(ctx context.Context, posts []*entity.Post) ([]*entity.Post, error) {
	postsMap := make(map[int64]*entity.Post, len(posts))
	postIDs := make([]int64, len(posts))
	for i, p := range posts {
		postsMap[p.ID] = p
		postIDs[i] = p.ID
	}
	cats, err := s.categoriesRepository.GetPostCategoriesByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, err
	}
	for _, cat := range cats {
		postsMap[cat.PostID].Categories = cat.Categories
	}
	return posts, nil
}
