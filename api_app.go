package identity

import (
	"context"
	"encoding/base64"
	"fmt"
	resty "github.com/go-resty/resty/v2"
)

// AppRequest IdentityRequest
type AppRequest struct {
	AuthValue string `json:"auth"`
	ServerURL string `json:"host"`
}

// NewAppRequestBasic NewAppRequestBasic
func NewAppRequestBasic(authKey, authSecret, serverURL string) *AppRequest {
	ks := authKey + ":" + authSecret
	ksStr := base64.StdEncoding.EncodeToString([]byte(ks))

	appReq := &AppRequest{
		AuthValue: fmt.Sprintf("Basic %s", ksStr),
		ServerURL: serverURL,
	}

	return appReq
}

// NewAppRequestJwt NewAppRequestJwt
func NewAppRequestJwt(token, serverURL string) *AppRequest {
	id := &AppRequest{
		AuthValue: fmt.Sprintf("Bearer %s", token),
		ServerURL:     serverURL,
	}

	return id
}

// GetApp GetApp
func (ir AppRequest) GetApp(ctx context.Context) (*Application, *AppError) {
	var res Application

	url := fmt.Sprintf("%s/v1/app", ir.ServerURL)

	// Request
	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}



// ============ private ============= //
// ============ private ============= //

func (ir AppRequest) getRequest(ctx context.Context) *resty.Request {
	return NewRequest(ctx).
		SetHeader("Authorization", ir.AuthValue).
		SetHeader(RequestIDKey, GenRequestID(ctx))
}

