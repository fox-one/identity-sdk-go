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
