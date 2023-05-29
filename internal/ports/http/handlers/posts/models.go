package posts

import (
	"time"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
)

func newPostsResponse(posts []*entity.Post) []*postResponse {
	out := make([]*postResponse, len(posts))
	for i, p := range posts {
		out[i] = newPostResponse(p, nil)
	}
	return out
}

func newPostResponse(post *entity.Post, comments []*entity.Comment) *postResponse {
	return &postResponse{
		ID:            post.ID,
		Title:         post.Title,
		AuthorID:      post.AuthorID,
		CreatedAt:     post.CreateAt,
		Categories:    newCategoriesResponse(post.Categories),
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
		Comments:      newCommentsResponse(comments),
	}
}

type postResponse struct {
	ID            int64               `json:"id"`
	Title         string              `json:"title"`
	AuthorID      int64               `json:"author_id"`
	CreatedAt     time.Time           `json:"created_at"`
	Categories    []*categoryResponse `json:"categories"`
	Content       string              `json:"content"`
	LikesCount    int64               `json:"likes_count"`
	DislikesCount int64               `json:"dislikes_count"`
	Comments      []*commentResponse  `json:"comments,omitempty"`
}

func newCategoriesResponse(categories []*entity.Category) []*categoryResponse {
	out := make([]*categoryResponse, len(categories))
	for i, cat := range categories {
		out[i] = newCategoryResponse(cat)
	}
	return out
}

func newCategoryResponse(category *entity.Category) *categoryResponse {
	return &categoryResponse{
		ID:    category.ID,
		Title: category.Title,
	}
}

type categoryResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func newCommentsResponse(comments []*entity.Comment) []*commentResponse {
	out := make([]*commentResponse, len(comments))
	for i, c := range comments {
		out[i] = newCommentResponse(c)
	}
	return out
}

func newCommentResponse(comment *entity.Comment) *commentResponse {
	return &commentResponse{
		ID:            comment.ID,
		AuthorID:      comment.AuthorID,
		Content:       comment.Content,
		CreatedAt:     comment.CreateAt,
		LikesCount:    comment.LikesCount,
		DislikesCount: comment.DislikesCount,
	}

}

type commentResponse struct {
	ID            int64     `json:"id"`
	AuthorID      int64     `json:"author_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	LikesCount    int64     `json:"likes"`
	DislikesCount int64     `json:"dislikes"`
}
