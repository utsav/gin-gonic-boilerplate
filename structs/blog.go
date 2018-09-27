package structs

type BlogFields struct {
	UserID uint `json:"user_id" binding:"required"`
	UpdateBlogFields
}

type UpdateBlogFields struct {
	Title   string `gorm:"type:varchar(100)" json:"title" binding:"required"`
	Content string `gorm:"type:text" json:"content" binding:"required"`
}
