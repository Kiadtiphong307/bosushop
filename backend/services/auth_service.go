// สำหรับการทำงานกับฐานข้อมูล
package services

import (
	"backend/database"
	"backend/models"
	"backend/validation"
	

	"golang.org/x/crypto/bcrypt"

)

// สมัครสมาชิก
func RegisterUser(input validation.RegisterInput) (*models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashed),
		Role:     "user",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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
