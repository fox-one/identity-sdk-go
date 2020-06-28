package identity

import (
	"context"
	"fmt"
	"strings"
)

// GetAllUsers GetAllUsers
func (ir AppRequest) GetAllUsers(ctx context.Context) ([]*User, *AppError) {
	var users []*User

	if err := Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/users", ir.ServerURL), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser GetUser
func (ir AppRequest) GetUser(ctx context.Context, userID uint64, profile, mixinAuth, foxAuth bool) (*User, *AppError) {
	var resp User

	var expand = make([]string, 0)
	if profile {
		expand = append(expand, "profile")
	}
	if mixinAuth {
		expand = append(expand, "authorizations.mixin")
	}
	if foxAuth {
		expand = append(expand, "authorizations.foxone")
	}

	url := fmt.Sprintf("%s/v1/users/%v/?expand=%s", ir.ServerURL, userID, strings.Join(expand, ","))

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser CreateUser
func (ir AppRequest) CreateUser(ctx context.Context, req *CreateUserReq) (*User, *AppError) {
	var user User

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/users"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
