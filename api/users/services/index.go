package UserServices

import (
	"errors"
	"fmt"
	UserRepository "go-template/api/users/repository"
	"go-template/utils"
)

func SignUp(UserCode string, Password string, Firstname string, Lastname string, BranchCode string) error {
	count, err := UserRepository.CountWhereUsername(UserCode)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Username already has!")
	}
	var hash = utils.Encode(Password)
	err = UserRepository.AddUsers(UserCode, hash, Firstname, Lastname, BranchCode)
	if err != nil {
		return err
	}
	return nil
}

func SignIn(UserCode string, Password string) (*string, error) {
	result, err := UserRepository.GetUserSignInByUserCode(UserCode)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	fmt.Print("test2")
	if !utils.Compare(Password, result.Password) {
		return nil, errors.New("Password is not compare!")
	}
	tokenString, err := utils.GenerateJwt(result.QuestionUserID, result.UserCode, result.BranchCode)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
