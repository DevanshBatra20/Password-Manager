package exceptions

import "fmt"

type InvalidToken struct {
	Token string
}

func (i InvalidToken) Error() string {
	return fmt.Sprintf("Invalid Token: '%v'", i.Token)
}
