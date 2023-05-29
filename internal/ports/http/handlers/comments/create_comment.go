package comments

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var req createCommentRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	id, err := c.commentsService.CreateComment(r.Context(), req.convertToComment())
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusCreated, util.NewSuccessResponse(
		struct {
			ID int64 `json:"id"`
		}{ID: id},
	))
}

type createCommentRequest struct {
	PostID   int64  `json:"post_id"`
	AuthorID int64  `json:"-"`
	Content  string `json:"content"`
}

func (r *createCommentRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *createCommentRequest) GetParams(req *http.Request) error {
	userID, ok := req.Context().Value("user_id").(int64)
	if !ok {
		return errors.New("undefined user_id")
	}
	r.AuthorID = userID
	return nil
}

func (r *createCommentRequest) Validate() error {
	if !util.IsPositiveNumber(r.PostID) {
		return errors.New("post id must be positive")
	}
	if !util.IsPositiveNumber(r.AuthorID) {
		return errors.New("author id must be positive")
	}
	if util.IsEmptyString(r.Content) {
		return errors.New("content mustn't be empty")
	}
	return nil
}

func (c *createCommentRequest) convertToComment() *entity.Comment {
	return &entity.Comment{
		PostID:   c.PostID,
		AuthorID: c.PostID,
		Content:  c.Content,
	}
}
