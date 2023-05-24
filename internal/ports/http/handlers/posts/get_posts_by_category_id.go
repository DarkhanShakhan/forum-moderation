package posts

import (
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) getPostsByCategoryID(w http.ResponseWriter, r *http.Request) {
	var req getPostsByCategoryIDRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}

	posts, err := c.postsService.GetPostsByCategory(r.Context(), req.ID)
	if err != nil {
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostsResponse(posts)))
}

type getPostsByCategoryIDRequest struct {
	ID int64
}

func (r *getPostsByCategoryIDRequest) GetParams(req *http.Request) error {
	id, err := util.GetIDFromPath(getPostByCategoryIDPattern, req.URL.Path)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

func (r *getPostsByCategoryIDRequest) Validate() error {
	if !util.IsPositiveNumber(r.ID) {
		return errors.New("id must be positive")
	}
	return nil
}
