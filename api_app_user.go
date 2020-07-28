package identity

import (
	"context"
	"encoding/json"
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

// GetUserByPhone GetUser
func (ir AppRequest) GetUserByPhone(ctx context.Context, phoneCode, phoneNumber string) (*User, *AppError) {
	var resp BasePageResponse

	url := fmt.Sprintf("%s/v1/users?phone_code=%s&phone_number=%s&limit=1", ir.ServerURL, phoneCode, phoneNumber)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	if len(resp.Items) > 0 {
		userBt, err := json.Marshal(resp.Items[0])
		if err != nil {
			return nil, NewAppError(err.Error())
		}

		user := new(User)
		err2 := json.Unmarshal(userBt, user)
		if err2 != nil {
			return nil, NewAppError(err2.Error())
		}
		return user, nil
	}

	return nil, nil
}

// VerifyUserPassword VerifyUserPassword
func (ir AppRequest) VerifyUserPassword(ctx context.Context, userID uint64, password string) (*User, *AppError) {
	var resp User

	url := fmt.Sprintf("%s/v1/users/%v/password/verify", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "POST", url, map[string]string{"password": password}, &resp); err != nil {
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
