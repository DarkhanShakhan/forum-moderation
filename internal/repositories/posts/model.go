package posts

import "github.com/DarkhanShakhan/forum-moderation/internal/domain/entity"

type model struct {
	ID            uint64
	Title         string
	AuthorID      uint64
	Content       string
	LikesCount    uint64
	DislikesCount uint64
}

func (m *model) convert() *entity.Post {
	if m == nil {
		return nil
	}
	return &entity.Post{
		ID:            m.ID,
		Title:         m.Title,
		AuthorID:      m.AuthorID,
		Content:       m.Content,
		LikesCount:    m.LikesCount,
		DislikesCount: m.DislikesCount,
	}
}

type models []*model

func (mm models) convert() []*entity.Post {
	out := make([]*entity.Post, len(mm))
	for i, m := range mm {
		out[i] = m.convert()
	}
	return out
}
