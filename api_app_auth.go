package identity

import (
	"context"
	"fmt"
)

// AuthByMixin AuthByMixin
func (ir AppRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*User, *AppError) {
	var userResp User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/mixin"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// AuthByFoxone AuthByFoxone
func (ir AppRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*User, *AppError) {
	var userResp User

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/foxone"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// GetAuths GetAuths
func (ir AppRequest) GetAuths(ctx context.Context, provider string, offset, limit int) (*AuthorizationList, *AppError) {
	var auths AuthorizationList
	var url = fmt.Sprintf("%s/v1/app/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, limit, offset)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
		return nil, err
	}

	return &auths, nil
}

// GenToken GenToken
func (ir AppRequest) GenToken(ctx context.Context, req *TokenCreateRequest) (*Token, *AppError) {
	var tokenRes Token
	var url = fmt.Sprintf("%s/v1/app/users/%v/tokens", ir.ServerURL, req.Audience)

	if err := Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

// GetAuthByOAuthID GetAuthByOAuthID
func (ir AppRequest) GetAuthByOAuthID(ctx context.Context, provider AuthProviderTypeEnum, oauthID string) (*Authorization, *AppError) {
	var auth Authorization
	var url = fmt.Sprintf("%s/v1/app/%s/auths/%s", ir.ServerURL, provider, oauthID)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}


// GetAuthsByOAuthIDs GetAuthsByOAuthIDs
func (ir AppRequest) GetAuthsByOAuthIDs(ctx context.Context, provider AuthProviderTypeEnum, oauthIDs []string) ([]*Authorization, *AppError) {
	var auth []*Authorization
	var url = fmt.Sprintf("%s/v1/app/%s/auths/batch", ir.ServerURL, provider)

	if err := Execute(ir.getRequest(ctx), "POST", url, oauthIDs, &auth); err != nil {
		return nil, err
	}

	return auth, nil
}

// GenMfaPhoneCode GenMfaPhoneCode
func (ir AppRequest) GenMfaPhoneCode(ctx context.Context, authReq *PhoneCodeVerifyRequest) (string, *AppError) {
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
func (ir AppRequest) VerifyMfaPhoneCode(ctx context.Context, authReq *PhoneCodeVerifyRequest) (*User, *AppError) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone/verify"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// BindAuth BindAuth
func (ir AppRequest) BindAuth(ctx context.Context, req AuthBindingRequest) (*Authorization, *AppError) {
	var auth Authorization
	if err := Execute(ir.getRequest(ctx), "PUT", fmt.Sprintf("%s/v1/app/users/%v/auths/%s/bind", ir.ServerURL, req.UserID, req.Provider), req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UnbindAuth BindAuth
func (ir AppRequest) UnbindAuth(ctx context.Context, userID uint64, provider AuthProviderTypeEnum) *AppError {
	var result map[string]interface{}
	if err := Execute(ir.getRequest(ctx), "DELETE", fmt.Sprintf("%s/v1/app/users/%v/auths/%s/bind", ir.ServerURL, userID, provider), nil, &result); err != nil {
		return err
	}

	if result["result"] == nil || result["result"].(string) != "ok" {
		return NewAppError("result error!")
	}

	return nil
}
