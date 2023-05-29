package admin

import (
	"errors"
	"net/http"

	appErr "github.com/DarkhanShakhan/forum-moderation/internal/errors"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req deleteCategoryRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	err := c.categoriesService.DeleteCategory(r.Context(), req.ID)
	if err != nil {
		if errors.Is(err, appErr.ErrCategoryNotFound) {
			util.SendError(w, err, http.StatusNotFound)
			return
		}
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type deleteCategoryRequest struct {
	ID int64
}

func (r *deleteCategoryRequest) GetParams(req *http.Request) error {
	id, err := util.GetIDFromPath(deleteCategoryPattern, req.URL.Path)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

func (r *deleteCategoryRequest) Validate() error {
	if !util.IsPositiveNumber(r.ID) {
		return errors.New("id must be positive")
	}
	return nil
}
