package identity

import (
	"context"
	"fmt"

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
func (r UserRequest) GetMe(ctx context.Context) (*UserAuthsResponse, error) {
	var res UserAuthsResponse
	err := httputils.Execute(r.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/user?expand=profile,authorizations.mixin,authorizations.foxone", r.ServerURL), nil, &res)
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
