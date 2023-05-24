package posts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var req createPostRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}

	id, err := c.postsService.CreatePost(r.Context(), req.convertToPost())
	if err != nil {
		//FIXME:
		return
	}
	util.SendData(w, http.StatusCreated, util.NewSuccessResponse(
		struct {
			ID int64 `json:"id"`
		}{ID: id},
	))
}

type createPostRequest struct {
	Title      string  `json:"title"`
	AuthorID   int64   `json:"-"`
	Content    string  `json:"content"`
	Categories []int64 `json:"categories"`
}

func (r *createPostRequest) convertToPost() *entity.Post {
	return &entity.Post{
		Title:      r.Title,
		AuthorID:   r.AuthorID,
		Categories: convertToCategories(r.Categories),
		Content:    r.Content,
	}
}

func convertToCategories(categories []int64) []*entity.Category {
	out := make([]*entity.Category, len(categories))
	for ix, id := range categories {
		out[ix].ID = id
	}
	return out
}

func (r *createPostRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *createPostRequest) GetParams(req *http.Request) error {
	userID, ok := req.Context().Value("user_id").(int64)
	if !ok {
		return errors.New("user_id isn't valid")
	}
	r.AuthorID = userID
	return nil
}

func (r *createPostRequest) Validate() error {
	if util.IsEmptyString(r.Title) {
		return errors.New("empty title")
	}
	if len(r.Content) > 0 && util.IsEmptyString(r.Content) {
		return errors.New("content with spaces")
	}
	if r.AuthorID <= 0 {
		return errors.New("empty author_id")
	}
	if util.IsEmptySlice(r.Categories) || !util.IsUniqueSlice(r.Categories) {
		return errors.New("empty or not unique category list")
	}
	return nil
}
