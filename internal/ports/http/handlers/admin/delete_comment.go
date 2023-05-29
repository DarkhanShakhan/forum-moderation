package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"
	appErr "github.com/DarkhanShakhan/forum-moderation/internal/errors"
	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	var req deleteCommentRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	err := c.commentsService.DeleteComment(r.Context(), req.ID, req.DeleteCategory, req.DeleteMessage)
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

type deleteCommentRequest struct {
	ID             int64
	DeletedBy      int64
	DeleteCategory string  `json:"delete_category"`
	DeleteMessage  *string `json:"delete_message,omitempty"`
}

func (r *deleteCommentRequest) GetParams(req *http.Request) error {
	id, err := util.GetIDFromPath(deleteCommentPattern, req.URL.Path)
	if err != nil {
		return err
	}
	r.ID = id
	userID, ok := req.Context().Value("user_id").(int64)
	if !ok {
		return errors.New("undefined user_id")
	}
	r.DeletedBy = userID
	return nil
}

func (r *deleteCommentRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *deleteCommentRequest) Validate() error {
	if !util.IsPositiveNumber(r.ID) {
		return errors.New("id must be positive")
	}
	if !util.IsPositiveNumber(r.DeletedBy) {
		return errors.New("user id must be positive")
	}
	if _, ok := enum.ParseStringToReportCategory(r.DeleteCategory); !ok {
		return errors.New("undefined delete category")
	}
	return nil
}
