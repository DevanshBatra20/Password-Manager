package exception

import "fmt"

type UserNotFound struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
}

func (u UserNotFound) Error() string {
	return fmt.Sprintf("User with Email:'%v' not found!", u.Email)
}
