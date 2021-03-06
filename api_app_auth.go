package identity

import (
	"context"
	"fmt"
)

// AuthByMixin AuthByMixin
func (ir AppRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*User, error) {
	var userResp User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/mixin"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// AuthByFoxone AuthByFoxone
func (ir AppRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*User, error) {
	var userResp User

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/foxone"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// AuthByFoxone AuthByFoxone
func (ir AppRequest) AuthByWechat(ctx context.Context, authReq *WechatAuthReq) (*User, error) {
	var userResp User

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/wechat"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// GetAuths GetAuths
func (ir AppRequest) GetAuths(ctx context.Context, provider string, offset, limit int) (*AuthorizationList, error) {
	var auths AuthorizationList
	var url = fmt.Sprintf("%s/v1/app/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, limit, offset)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
		return nil, err
	}

	return &auths, nil
}

// GenToken GenToken
func (ir AppRequest) GenToken(ctx context.Context, req *TokenCreateRequest) (*Token, error) {
	var tokenRes Token
	var url = fmt.Sprintf("%s/v1/app/users/%v/tokens", ir.ServerURL, req.Audience)

	if err := Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

// RemoveToken RemoveToken
func (ir AppRequest) RemoveToken(ctx context.Context, userID uint64) error {
	var result interface{}
	var url = fmt.Sprintf("%s/v1/app/users/%v/tokens", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "DELETE", url, nil, &result); err != nil {
		return err
	}

	return nil
}

// GenMfaToken GenMfaToken
func (ir AppRequest) GenMfaToken(ctx context.Context, userID uint64) (string, error) {
	var tokenRes struct {
		MfaToken string `json:"mfa_token"`
	}
	var url = fmt.Sprintf("%s/v1/app/users/%v/mfa_tokens", ir.ServerURL, userID)

	if err := Execute(ir.getRequest(ctx), "POST", url, nil, &tokenRes); err != nil {
		return "", err
	}

	return tokenRes.MfaToken, nil
}

// GetAuthByOAuthID GetAuthByOAuthID
func (ir AppRequest) GetAuthByOAuthID(ctx context.Context, provider AuthProviderTypeEnum, oauthID string) (*Authorization, error) {
	var auth Authorization
	var url = fmt.Sprintf("%s/v1/app/%s/auths/%s", ir.ServerURL, provider, oauthID)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// GetAuthsByOAuthIDs GetAuthsByOAuthIDs
func (ir AppRequest) GetAuthsByOAuthIDs(ctx context.Context, provider AuthProviderTypeEnum, oauthIDs []string) ([]*Authorization, error) {
	var auth []*Authorization
	var url = fmt.Sprintf("%s/v1/app/%s/auths/batch", ir.ServerURL, provider)

	if err := Execute(ir.getRequest(ctx), "POST", url, oauthIDs, &auth); err != nil {
		return nil, err
	}

	return auth, nil
}

// GenMfaPhoneCode GenMfaPhoneCode
func (ir AppRequest) GenMfaPhoneCode(ctx context.Context, authReq *PhoneCodeVerifyRequest) (string, error) {
	var result map[string]interface{}
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone"), authReq, &result); err != nil {
		return "", err
	}

	if result["code"] == nil {
		return "", NewAppError("result error!")
	}

	return result["code"].(string), nil
}

// VerifyMfaPhoneCode VerifyMfaPhoneCode
func (ir AppRequest) VerifyMfaPhoneCode(ctx context.Context, verifyReq *PhoneCodeVerifyRequest) (*User, error) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone/verify"), verifyReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// BindAuth BindAuth
func (ir AppRequest) BindAuth(ctx context.Context, req AuthBindingRequest) (*Authorization, error) {
	var auth Authorization
	if err := Execute(ir.getRequest(ctx), "PUT", fmt.Sprintf("%s/v1/app/users/%v/auths/%s/bind", ir.ServerURL, req.UserID, req.Provider), req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UnbindAuth BindAuth
func (ir AppRequest) UnbindAuth(ctx context.Context, userID uint64, provider AuthProviderTypeEnum) error {
	var result map[string]interface{}
	if err := Execute(ir.getRequest(ctx), "DELETE", fmt.Sprintf("%s/v1/app/users/%v/auths/%s/bind", ir.ServerURL, userID, provider), nil, &result); err != nil {
		return err
	}

	if result["result"] == nil || result["result"].(string) != "ok" {
		return NewAppError("result error!")
	}

	return nil
}

// GetAuthorizedKey GetAuthorizedKey
func (ir AppRequest) GetAuthorizedKey(ctx context.Context, authKey string) (*AuthorizedKey, error) {
	var res AuthorizedKey

	url := fmt.Sprintf("%s/v1/auth_keys/%s", ir.ServerURL, authKey)

	// Request
	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}
