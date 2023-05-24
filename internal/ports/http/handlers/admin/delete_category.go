package admin

import (
	"errors"
	"net/http"

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
		//FIXME:
		return
	}
	util.SendData(w, http.StatusOK, util.NewEmptySuccessResponse())
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
