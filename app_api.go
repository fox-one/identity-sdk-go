package identity

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"

	httpUtil "github.com/fox-one/identity-sdk-go/utils"
)

// IDRequest IdentityRequest
type IDRequest struct {
	AuthValue string `json:"auth"`
	ServerURL string `json:"host"`
}

// NewIDRequestBasic NewIDRequestBasic
func NewIDRequestBasic(authKey, authSecret, serverURL string) *IDRequest {
	ks := authKey + ":" + authSecret
	ksStr := base64.StdEncoding.EncodeToString([]byte(ks))

	id := &IDRequest{
		AuthValue: fmt.Sprintf("Basic %s", ksStr),
		ServerURL: serverURL,
	}
	return id
}

func (ir IDRequest) getRequest(ctx context.Context) *resty.Request {
	return httpUtil.NewRequest(ctx).
		SetHeader("Authorization", ir.AuthValue).
		SetHeader(httpUtil.RequestIDKey, httpUtil.GenRequestID(ctx))
}

// GetAllUsers GetAllUsers
func (ir IDRequest) GetAllUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	if err := httpUtil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/users", ir.ServerURL), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser GetUser
func (ir IDRequest) GetUser(ctx context.Context, userID uint64) (*User, error) {
	var user User
	if err := httpUtil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s%s%d", ir.ServerURL, "/v1/users/", userID), nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser CreateUser
func (ir IDRequest) CreateUser(ctx context.Context, req *CreateUserReq) (*User, error) {
	var user User
	if err := httpUtil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/users"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
