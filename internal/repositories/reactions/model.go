package reactions

import (
	"github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"
)

type model struct {
	EntityID int64
	UserID   int64
	Like     bool
}

func (m *model) convertToReaction() *entity.Reaction {
	return &entity.Reaction{
		EntityID: m.EntityID,
		UserID:   m.UserID,
		Like:     m.Like,
	}
}

type models []*model

func (mm models) convertToReactions() []*entity.Reaction {
	out := make([]*entity.Reaction, len(mm))
	for i, m := range mm {
		out[i] = m.convertToReaction()
	}
	return out
}
