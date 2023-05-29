package comments

import "github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"

type model struct {
	ID            int64
	PostID        int64
	AuthorID      int64
	Content       string
	LikesCount    int64
	DislikesCount int64
}

func (m *model) convert() *entity.Comment {
	if m == nil {
		return nil
	}
	return &entity.Comment{
		ID:            m.ID,
		PostID:        m.PostID,
		Content:       m.Content,
		LikesCount:    m.LikesCount,
		DislikesCount: m.DislikesCount,
	}
}

type models []*model

func (mm models) convert() []*entity.Comment {
	out := make([]*entity.Comment, len(mm))
	for i, m := range mm {
		out[i] = m.convert()
	}
	return out
}
