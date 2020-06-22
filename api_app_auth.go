package identity

import (
	"context"
	"fmt"

	httputil "github.com/fox-one/identity-sdk-go/utils"
)

// AuthByMixin AuthByMixin
func (ir IDRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*UserAuthsResponse, error) {
	var userResp UserAuthsResponse

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/mixin"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// AuthByFoxone AuthByFoxone
func (ir IDRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*UserAuthsResponse, error) {
	var userResp UserAuthsResponse

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/foxone"), authReq, &userResp); err != nil {
		return nil, err
	}

	return &userResp, nil
}

// GetAuths GetAuths
func (ir IDRequest) GetAuths(ctx context.Context, provider string, start, limit int) ([]*AuthorizationResponse, error) {
	var auths []*AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/app/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, start, limit)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
		return nil, err
	}

	return auths, nil
}
