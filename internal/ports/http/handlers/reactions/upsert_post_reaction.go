package reactions

import (
	"net/http"

	"github.com/DarkhanShakhan/forum-moderation/internal/util"
)

func (c *controller) upsertPostReactions(w http.ResponseWriter, r *http.Request) {
	var req upsertReactionRequest
	if err := util.GetData(r, &req); err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	if err := c.reactionsService.UpsertPostReaction(r.Context(), req.convertToReaction()); err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted) //FIXME: not sure
}
