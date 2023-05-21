package posts

import "github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"

func newPostsResponse(posts []*entity.Post) []*postResponse {
	out := make([]*postResponse, len(posts))
	for i, p := range posts {
		out[i] = newPostResponse(p)
	}
	return out
}

func newPostResponse(post *entity.Post) *postResponse {
	return &postResponse{
		ID:            post.ID,
		Title:         post.Title,
		AuthorID:      post.AuthorID,
		Categories:    newCategoriesResponse(post.Categories),
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
	}
}

type postResponse struct {
	ID            int64               `json:"id"`
	Title         string              `json:"title"`
	AuthorID      int64               `json:"author_id"`
	Categories    []*categoryResponse `json:"categories"`
	Content       string              `json:"content"`
	LikesCount    int64               `json:"likes_count"`
	DislikesCount int64               `json:"dislikes_count"`
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
