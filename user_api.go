package identity

import (
	"context"
	"fmt"

	"github.com/fox-one/identity-sdk-go/utils"
	"github.com/go-resty/resty/v2"
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

func (r UserRequest) getRequest(ctx context.Context) *resty.Request {
	return utils.NewRequest(ctx).
		SetHeader("Authorization", r.Authorization).
		SetHeader(utils.RequestIDKey, utils.GenRequestID(ctx))
}

// FetchMe FetchMe
func (r UserRequest) FetchMe(ctx context.Context) (*UserAuthsResponse, error) {
	var res UserAuthsResponse
	err := utils.Execute(r.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/user?expand=profile,authorizations.mixin,authorizations.foxone", r.ServerURL), nil, &res)
	if nil != err {
		return nil, err
	}
	return &res, nil
}
