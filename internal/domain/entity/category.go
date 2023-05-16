package entity

type Category struct {
	ID    int64
	Title string
}

type PostCategories struct {
	PostID     int64
	Categories []*Category
}
