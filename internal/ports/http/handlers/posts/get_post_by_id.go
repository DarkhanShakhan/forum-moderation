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
	ctx := r.Context()
	post, err := c.postsService.GetPostByID(ctx, req.ID)
	if err != nil {
		util.SendError(w, err, http.StatusNotFound) //FIXME: check for error
		return
	}
	// TODO: async
	comments, err := c.commentsService.GetCommentsByPostID(ctx, post.ID)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostResponse(post, comments)))
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
