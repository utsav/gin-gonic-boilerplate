package structs

type UserFields struct {
	Email    string `gorm:"type:varchar(100);unique" json:"email" binding:"required,email"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty" binding:"required"`
	UpdateUserFields
}

type UpdateUserFields struct {
	Name string `gorm:"type:varchar(100)" json:"name" binding:"required"`
}
