package identity

import (
	"context"
	"encoding/base64"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// IDRequest IdentityRequest
type IDRequest struct {
	AuthKey    string `json:"key"`
	AuthSecret string `json:"secret"`
	ServerURL  string `json:"host"`
}

func (ir IDRequest) getRequest(ctx context.Context) *resty.Request {
	ks := ir.AuthKey + ":" + ir.AuthSecret
	ksStr := base64.StdEncoding.EncodeToString([]byte(ks))
	fmt.Println("identity.Authorization:", ksStr)

	return Request(ctx).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Charset":       "utf-8",
			"Authorization": fmt.Sprintf("Basic %s", ksStr),
		})
}

// GetAllUsers GetAllUsers
func (ir IDRequest) GetAllUsers(ctx context.Context) ([]*User, error) {
	resp, err := ir.getRequest(ctx).Get(ir.ServerURL + "/v1/users")
	if err != nil {
		return nil, err
	}

	var users []*User
	err = ParseResponse(resp, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser GetUser
func (ir IDRequest) GetUser(ctx context.Context, userID uint64) (*User, error) {
	url := fmt.Sprintf("%s%s%d", ir.ServerURL, "/v1/users/", userID)
	resp, err := ir.getRequest(ctx).Get(url)
	if err != nil {
		return nil, err
	}

	var user = new(User)
	err = ParseResponse(resp, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser CreateUser
func (ir IDRequest) CreateUser(ctx context.Context, req *CreateUserReq) (*User, error) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/users"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// AuthByMixin AuthByMixin
func (ir IDRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*User, error) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/auths/mixin"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// AuthByFoxone AuthByFoxone
func (ir IDRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*User, error) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/auths/foxone"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Login Login
func (ir IDRequest) Login(ctx context.Context, req *LoginRequest) (*User, error) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/login"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
