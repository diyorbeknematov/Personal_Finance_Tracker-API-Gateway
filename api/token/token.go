package token

import (
	"fmt"
)

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func TokenClaimsParse(val any) (*Claims, error) {
	claims, ok := val.(*Claims)
	if !ok {
		return nil, fmt.Errorf("cannot parse token claims")
	}

	return claims, nil
}

func (tc *Claims) GetId() string {
	return tc.Id
}

func (tc *Claims) GetEmail() string {
	return tc.Email
}

func (tc *Claims) GetRole() string {
	return tc.Role
}
