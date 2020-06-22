package identity

import (
	"context"
	"fmt"

	httputil "github.com/fox-one/identity-sdk-go/utils"
)

// AuthByMixin AuthByMixin
func (ir IDRequest) AuthByMixin(ctx context.Context, authReq *MixinAuthReq) (*User, error) {
	var user User
	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/mixin"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// AuthByFoxone AuthByFoxone
func (ir IDRequest) AuthByFoxone(ctx context.Context, authReq *FoxoneAuthReq) (*User, error) {
	var user User
	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/app/auths/foxone"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAuths GetAuths
func (ir IDRequest) GetAuths(ctx context.Context, provider string, start, limit int) (*User, error) {
	var user User

	var url = fmt.Sprintf("%s/v1/app/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, start, limit)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
