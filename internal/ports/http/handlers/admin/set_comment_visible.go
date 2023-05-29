package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	appErr "github.com/DarkhanShakhan/forum-moderation/internal/errors"
	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) setCommentVisibleHandler(w http.ResponseWriter, r *http.Request) {
	var req setCommentVisibleRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	err := c.commentsService.SetVisible(r.Context(), req.CommentID, req.Visible)
	if err != nil {
		if errors.Is(err, appErr.ErrCommentNotFound) {
			util.SendError(w, err, http.StatusNotFound)
			return
		}
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type setCommentVisibleRequest struct {
	CommentID int64 `json:"comment_id"`
	Visible   bool  `json:"visible"`
}

func (r *setCommentVisibleRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *setCommentVisibleRequest) Validate() error {
	if !util.IsPositiveNumber(r.CommentID) {
		return errors.New("id must be positive")
	}
	// FIXME: check visible *bool
	return nil
}
