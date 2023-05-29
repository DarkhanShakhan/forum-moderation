package entity

type Comment struct {
	ID            int64
	PostID        int64
	AuthorID      int64
	Content       string
	LikesCount    int64
	DislikesCount int64
}
