package admin_model

type Admin struct {
	Username string `gorm:"primarykey" json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Getadmin(user, pass string) (bool, error) {
	return true, nil
}
