package exception

import "fmt"

type InvalidCredentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (i InvalidCredentials) Error() string {
	return fmt.Sprintf("Invalid Password:'%v' for Email:'%v'", i.Password, i.Email)
}
