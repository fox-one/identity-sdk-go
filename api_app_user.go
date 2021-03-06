package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// GetAllUsers GetAllUsers
func (ir AppRequest) GetAllUsers(ctx context.Context) ([]*User, error) {
	var users []*User

	if err := Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/users", ir.ServerURL), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser GetUser
func (ir AppRequest) GetUser(ctx context.Context, userID uint64, profile, mixinAuth, foxAuth bool, wechatAuth bool, queries ...string) (*User, error) {
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
	if wechatAuth {
		expand = append(expand, "authorizations.wechat")
	}

	url := fmt.Sprintf("%s/v1/users/%v/?expand=%s&%s", ir.ServerURL, userID, strings.Join(expand, ","), strings.Join(queries, "&"))

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// BatchGetUsers BatchGetUsers
func (ir AppRequest) BatchGetUsers(ctx context.Context, userIDs []uint64, profile, mixinAuth, foxAuth, wechatAuth bool, query ...string) ([]*User, error) {
	var resp []*User

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
	if wechatAuth {
		expand = append(expand, "authorizations.wechat")
	}

	userIDStr := make([]string, 0)
	for _, id := range userIDs {
		userIDStr = append(userIDStr, fmt.Sprintf("%v", id))
	}

	url := fmt.Sprintf("%s/v1/users?id=%s&expand=%s", ir.ServerURL, strings.Join(userIDStr, ","), strings.Join(expand, ","))

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUserByPhone GetUser
func (ir AppRequest) GetUserByPhone(ctx context.Context, phoneCode, phoneNumber string) (*User, error) {
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
func (ir AppRequest) VerifyUserPassword(ctx context.Context, userID uint64, password string) (*User, error) {
	var resp User

	url := fmt.Sprintf("%s/v1/users/%v/password/verify", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "POST", url, map[string]string{"password": password}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser CreateUser
func (ir AppRequest) CreateUser(ctx context.Context, req *CreateUserReq) (*User, error) {
	var user User

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/users"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// SetPassword SetPassword
func (ir AppRequest) SetPassword(ctx context.Context, userID uint64, password string) (*User, error) {
	var user User

	url := fmt.Sprintf("%s/v1/users/%v/password", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "PUT", url, map[string]string{"password": password}, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ChangePassword ChangePassword
func (ir AppRequest) ChangePassword(ctx context.Context, userID uint64, oldPassword, newPassword string) (*User, error) {
	var user User

	url := fmt.Sprintf("%s/v1/users/%v/password", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "PATCH", url, map[string]string{"new_password": newPassword, "old_password": oldPassword}, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ChangePhone ChangePhone
func (ir AppRequest) ChangePhone(ctx context.Context, req *UserModifyReq) (*User, error) {
	var user User

	url := fmt.Sprintf("%s/v1/users/%d/phone", ir.ServerURL, req.UserID)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateMfa UpdateMfa
func (ir AppRequest) UpdateMfa(ctx context.Context, req *MfaCredentialRequest) (*MfaCredential, error) {
	var resp MfaCredential

	url := fmt.Sprintf("%s/v1/users/%d/mfa/two_factor", ir.ServerURL, req.UserID)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMfa UpdateMfa
func (ir AppRequest) RemoveMfa(ctx context.Context, userID uint64) (bool, error) {
	var result struct {
		Success bool `json:"success"`
	}

	url := fmt.Sprintf("%s/v1/users/%d/mfa/two_factor", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "DELETE", url, nil, &result); err != nil {
		return false, err
	}

	return result.Success, nil
}
