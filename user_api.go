package identity

import (
	"context"
	"fmt"
	"strings"

	httputils "github.com/fox-one/identity-sdk-go/utils"
	resty "github.com/go-resty/resty/v2"
)

// UserRequest UserRequest
type UserRequest struct {
	Authorization string `json:"token"`
	ServerURL     string `json:"server_url"`
}

// NewUserRequestJwt NewUserRequestJwt
func NewUserRequestJwt(token, serverURL string) *UserRequest {
	id := &UserRequest{
		Authorization: fmt.Sprintf("Bearer %s", token),
		ServerURL:     serverURL,
	}
	return id
}

// GetMe GetMe
func (r UserRequest) GetMe(ctx context.Context, userID uint64, profile, mixinAuth, foxAuth bool) (*UserAuthsResponse, error) {
	var res UserAuthsResponse

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

	err := httputils.Execute(r.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/user?expand=%s", r.ServerURL, strings.Join(expand, ",")), nil, &res)
	if nil != err {
		return nil, err
	}
	return &res, nil
}

// ============ private ============= //
// ============ private ============= //

func (r UserRequest) getRequest(ctx context.Context) *resty.Request {
	return httputils.NewRequest(ctx).
		SetHeader("Authorization", r.Authorization).
		SetHeader(httputils.RequestIDKey, httputils.GenRequestID(ctx))
}
