package user_model

type User struct {
	Username string `gorm:"primarykey" json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
