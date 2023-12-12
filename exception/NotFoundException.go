package exception

import "fmt"

type UserNotFound struct {
	Email string `json:"email"`
}

func (u UserNotFound) Error() string {
	return fmt.Sprintf("User with Email:'%v' not found!", u.Email)
}
