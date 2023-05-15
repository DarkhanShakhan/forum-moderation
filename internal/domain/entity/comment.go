package entity

type Comment struct {
	ID            string
	PostID        string
	AuthorID      string
	Content       string
	LikesCount    uint64
	DislikesCount uint64
}
