package helper

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateLogin(email, password string) error{
	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}
	if strings.Contains(password,"'") || strings.Contains(password,";") {
		panic("You are under arrest! Attempted SQL injection detected.")
	}

	//authentication process
	if password=="123456"{
		return nil // login success
	}else{
		return errors.New("invalid password")
	}
}

//checking email format valid
func isValidEmail(email string) bool{
	re:=regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
