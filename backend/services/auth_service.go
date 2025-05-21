// สำหรับการทำงานกับฐานข้อมูล
package services

import (
	"backend/database"
	"backend/models"
	

	"golang.org/x/crypto/bcrypt"

)

// สมัครสมาชิก
func RegisterUser(user *models.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashed)
	return database.DB.Create(user).Error
}

// เข้าสู่ระบบ
func Authenticate(email, password string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
