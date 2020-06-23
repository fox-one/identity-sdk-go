package identity

import (
	"context"
	"fmt"

	httputil "github.com/fox-one/identity-sdk-go/utils"
)

// AuthByMixin AuthByMixin
func (ir IDRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*User, error) {
	var userResp User

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/mixin"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// AuthByFoxone AuthByFoxone
func (ir IDRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*User, error) {
	var userResp User

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/foxone"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// GetAuths GetAuths
func (ir IDRequest) GetAuths(ctx context.Context, provider string, offset, limit int) (*AuthorizationList, error) {
	var auths AuthorizationList
	var url = fmt.Sprintf("%s/v1/app/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, limit, offset)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
		return nil, err
	}

	return &auths, nil
}

// GenToken GenToken
func (ir IDRequest) GenToken(ctx context.Context, req *TokenCreateRequest) (*Token, error) {
	var tokenRes Token
	var url = fmt.Sprintf("%s/v1/app/users/%v/tokens", ir.ServerURL, req.Audience)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}
