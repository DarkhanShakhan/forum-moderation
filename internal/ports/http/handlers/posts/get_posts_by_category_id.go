package posts

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) getPostsByCategoryID(w http.ResponseWriter, r *http.Request) {
	var req getPostsByCategoryIDRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}

	postsByCat, err := c.postsService.GetPostsByCategory(r.Context(), req.ID)
	if err != nil {
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(getPostsByCategoryResponse{
		Category: newCategoryResponse(postsByCat.Category),
		Posts:    newPostsResponse(postsByCat.Posts),
	}))
}

type getPostsByCategoryIDRequest struct {
	ID int64
}

func (r *getPostsByCategoryIDRequest) GetParams(req *http.Request) error {
	//FIXME: regex
	path := strings.Split(req.URL.Path, "/")
	id, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		return err
	}
	r.ID = int64(id)
	return nil
}

type getPostsByCategoryResponse struct {
	Category *categoryResponse `json:"category"`
	Posts    []*postResponse   `json:"posts"`
}
