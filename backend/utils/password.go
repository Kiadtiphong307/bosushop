package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword รับ string แล้วแปลงเป็น hash
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashed)
}

// CheckPassword เปรียบเทียบ plain กับ hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
