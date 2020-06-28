package identity

import (
	"context"
	"fmt"
	"strings"

	resty "github.com/go-resty/resty/v2"
)

// UserRequest UserRequest
type UserRequest struct {
	Authorization string `json:"token"`
	ServerURL     string `json:"server_url"`
}

// NewUserRequestJwt NewUserRequestJwt
func NewUserRequestJwt(token, serverURL string) *UserRequest {
	userReq := &UserRequest{
		Authorization: fmt.Sprintf("Bearer %s", token),
		ServerURL:     serverURL,
	}

	return userReq
}

// GetMe GetMe
func (r UserRequest) GetMe(ctx context.Context, profile, mixinAuth, foxAuth bool) (*User, *AppError) {
	var res User

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

	url := fmt.Sprintf("%s/v1/user?expand=%s", r.ServerURL, strings.Join(expand, ","))

	// Request
	if err := Execute(r.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}

// ============ private ============= //
// ============ private ============= //

func (r UserRequest) getRequest(ctx context.Context) *resty.Request {
	return NewRequest(ctx).
		SetHeader("Authorization", r.Authorization).
		SetHeader(RequestIDKey, GenRequestID(ctx))
}
