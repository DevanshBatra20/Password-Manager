package exception

import "fmt"

type InvalidToken struct {
	Token string `json:"token"`
}

func (i InvalidToken) Error() string {
	return fmt.Sprintf("Invalid Token:'%v'", i.Token)
}
