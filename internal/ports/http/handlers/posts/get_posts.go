package posts

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := c.postsService.GetPosts(r.Context())
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusOK, util.NewSuccessResponse(newPostsResponse(posts)))
}
