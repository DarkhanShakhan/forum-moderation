package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req createCategoryRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	id, err := c.categoriesService.CreateCategory(r.Context(), req.Title)
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

type createCategoryRequest struct {
	Title string `json:"title"`
}

func (r *createCategoryRequest) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *createCategoryRequest) Validate() error {
	if util.IsEmptyString(r.Title) {
		return errors.New("title must not be empty")
	}
	return nil
}
