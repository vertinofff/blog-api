package dto

type CreatePost struct {
	Title   string `json:"title" binding:"required,min=5,max=256"`
	Content string `json:"content" binding:"required,min=6,max=100000"`
	// AuthorId remains decodable for API compatibility but is ignored.
	AuthorId uint `json:"authorId"`
}

type UpdatePost struct {
	Title   string `json:"title" binding:"required,min=5,max=256"`
	Content string `json:"content" binding:"required,min=6,max=100000"`
}
