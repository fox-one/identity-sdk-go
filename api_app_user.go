package identity

import (
	"context"
	"fmt"
	"strings"

	httputil "github.com/fox-one/identity-sdk-go/utils"
)

// GetAllUsers GetAllUsers
func (ir IDRequest) GetAllUsers(ctx context.Context) ([]*UserAuthsResponse, error) {
	var users []*UserAuthsResponse

	if err := httputil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/users", ir.ServerURL), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser GetUser
func (ir IDRequest) GetUser(ctx context.Context, userID uint64, profile, mixinAuth, foxAuth bool) (*UserAuthsResponse, error) {
	var resp UserAuthsResponse

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

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser CreateUser
func (ir IDRequest) CreateUser(ctx context.Context, req *CreateUserReq) (*UserAuthsResponse, error) {
	var user UserAuthsResponse

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/users"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
