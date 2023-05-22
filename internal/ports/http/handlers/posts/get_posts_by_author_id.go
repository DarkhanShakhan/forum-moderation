package posts

import (
	"errors"
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) getPostsByAuthorID(w http.ResponseWriter, r *http.Request) {
	var req getPostsByAuthorIDRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	posts, err := c.postsService.GetPostsByAuthorID(r.Context(), req.ID)
	if err != nil {
		//TODO: deal with errors
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostsResponse(posts)))
}

type getPostsByAuthorIDRequest struct {
	ID int64
}

func (r *getPostsByAuthorIDRequest) GetParams(req *http.Request) error {
	ctx := req.Context()
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return errors.New("invalid user id")
	}
	r.ID = userID
	return nil
}

func (r *getPostsByAuthorIDRequest) Validate() error {
	if r.ID <= 0 {
		return errors.New("id must be positive number")
	}
	return nil
}
