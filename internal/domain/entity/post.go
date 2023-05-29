package entity

import "time"

type Post struct {
	ID            int64
	Title         string
	AuthorID      int64
	CreateAt      time.Time
	Categories    []*Category
	Content       string
	LikesCount    int64
	DislikesCount int64
}
