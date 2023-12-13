package exception

import "fmt"

type UserAlreadyExists struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u UserAlreadyExists) Error() string {
	return fmt.Sprintf("User with email: %v or phone: %v already exists", u.Email, u.Phone)
}
