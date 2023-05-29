package entity

import "time"

type Comment struct {
	ID            int64
	PostID        int64
	AuthorID      int64
	CreateAt      time.Time
	Content       string
	LikesCount    int64
	DislikesCount int64
}
