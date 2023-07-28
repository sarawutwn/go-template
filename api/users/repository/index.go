package UserRepository

import (
	"errors"
	Schema "go-template/api/users/schema"
	"go-template/database"
	"time"

	"github.com/google/uuid"
)

func CountWhereUsername(user_code string) (int, error) {
	db := database.DB
	query := "SELECT COUNT(*) FROM question_users WHERE user_code=$1"
	var count int
	err := db.QueryRow(query, user_code).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserSignInByUserCode(user_code string) (*Schema.UserSignIn, error) {
	db := database.DB
	query := `
		SELECT question_user_id, user_code, branch_code, password FROM question_users
		WHERE user_code=$1
	`
	users := Schema.UserSignIn{}
	err := db.QueryRow(query, user_code).Scan(&users.QuestionUserID, &users.UserCode, &users.BranchCode, &users.Password)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func AddUsers(UserCode string, Password string, Firstname string, Lastname string, BranchCode string) error {
	db := database.DB
	query := `
		insert into question_users(question_user_id, user_code, password, firstname, lastname, branch_code, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	newUUID := uuid.New()
	result, err := db.Exec(query, newUUID.String(), UserCode, Password, Firstname, Lastname, BranchCode, time.Now(), time.Now())
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Cannot insert to database!")
	}
	return nil
}
