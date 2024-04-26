package exceptions

import "fmt"

type TokenExpired struct {
	Token string
}

func (e TokenExpired) Error() string {
	return fmt.Sprintf("Token: %v is expired", e.Token)
}
