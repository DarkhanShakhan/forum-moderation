package entity

type Post struct {
	ID            uint64
	Title         string
	AuthorID      uint64
	Categories    []Category
	Content       string
	LikesCount    uint64
	DislikesCount uint64
}
