package utils

import (
	"go-template/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(user_id string, user_code string, user_branch_code string) (string, error) {
	var privateKey = config.GetEnvConfig("SECRET_KEY")
	claims := jwt.MapClaims{
		"question_user_id": user_id,
		"user_code":        user_code,
		"user_branch_code": user_branch_code,
		"exp":              time.Now().Add(time.Hour * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "Fail", err
	}
	return t, nil
}
