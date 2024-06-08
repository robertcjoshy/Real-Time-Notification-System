package admin_model

import "fmt"

type Admin struct {
	Username string `gorm:"primarykey" json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Getadmin(user, pass string) (bool, error) {
	if user == "robertcjoshy" && pass == "12345678" {
		return true, nil
	}
	return false, fmt.Errorf("invalid admin data")
}
