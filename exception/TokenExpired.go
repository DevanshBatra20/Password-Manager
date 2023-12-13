package exception

import "fmt"

type TokenExpired struct {
	Token string `json:"token"`
}

func (e TokenExpired) Error() string {
	return fmt.Sprintf("Token: %v is expired", e.Token)
}
