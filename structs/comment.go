package structs

type CommentFields struct {
	BlogID uint `json:"blog_id" binding:"required"`
	CommentUpdateFields
}

type CommentUpdateFields struct {
	Content string `gorm:"type:text" json:"content" binding:"required"`
}
