package posts

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

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
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostResponse(post)))
}

type getPostByIDRequest struct {
	ID int64
}

func (r *getPostByIDRequest) GetParams(req *http.Request) error {
	//FIXME: regex
	path := strings.Split(req.URL.Path, "/")
	id, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		return err
	}
	r.ID = int64(id)
	return nil
}

func (r *getPostByIDRequest) Validate() error {
	if r.ID <= 0 {
		return errors.New("id must be positive number")
	}
	return nil
}
