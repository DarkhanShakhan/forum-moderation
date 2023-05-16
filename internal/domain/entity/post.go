package entity

type Post struct {
	ID            int64
	Title         string
	AuthorID      int64
	Categories    []*Category
	Content       string
	LikesCount    int64
	DislikesCount int64
}
