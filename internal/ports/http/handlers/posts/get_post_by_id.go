package posts

import (
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

// TODO: errors
func (c *controller) getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req getPostByIDRequest
	if err := util.GetData(r, &req); err != nil {
		return
	}
	post, err := c.postsService.GetPostByID(r.Context(), req.ID)
	if err != nil {
		util.SendError(w, err, http.StatusNotFound)
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostResponse(post)))
}

type getPostByIDRequest struct {
	ID int64
}

func (r *getPostByIDRequest) GetParams(req *http.Request) error {
	id, err := util.GetIDFromPath(getPostByIDPattern, req.URL.Path)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

func (r *getPostByIDRequest) Validate() error {
	if !util.IsPositiveNumber(r.ID) {
		return errors.New("id must be positive number")
	}
	return nil
}
