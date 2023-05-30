package reactions

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

type upsertReactionRequest struct {
	EntityID int64 `json:"entity_id"`
	UserID   int64 `json:"-"`
	Like     bool  `json:"like"`
}

func (r *upsertReactionRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *upsertReactionRequest) GetParams(req *http.Request) error {
	userID, ok := req.Context().Value("user_id").(int64)
	if !ok {
		return errors.New("user_id is not defined")
	}
	r.UserID = userID
	return nil
}

func (r *upsertReactionRequest) Validate() error {
	if !util.IsPositiveNumber(r.EntityID) {
		return errors.New("entity id must be positive number")
	}
	if !util.IsPositiveNumber(r.UserID) {
		return errors.New("user id must be positive number")
	}
	return nil
}

func (r *upsertReactionRequest) convertToReaction() *entity.Reaction {
	return &entity.Reaction{
		EntityID: r.EntityID,
		UserID:   r.UserID,
		Like:     r.Like,
	}
}
