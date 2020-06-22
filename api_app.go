package identity

import (
	"context"
	"encoding/base64"
	"fmt"

	httputil "github.com/fox-one/identity-sdk-go/utils"

	resty "github.com/go-resty/resty/v2"
)

// IDRequest IdentityRequest
type IDRequest struct {
	AuthValue string `json:"auth"`
	ServerURL string `json:"host"`
}

// NewIDRequestBasic NewIDRequestBasic
func NewIDRequestBasic(authKey, authSecret, serverURL string) *IDRequest {
	ks := authKey + ":" + authSecret
	ksStr := base64.StdEncoding.EncodeToString([]byte(ks))

	id := &IDRequest{
		AuthValue: fmt.Sprintf("Basic %s", ksStr),
		ServerURL: serverURL,
	}

	return id
}

// ============ private ============= //
// ============ private ============= //

func (ir IDRequest) getRequest(ctx context.Context) *resty.Request {
	return httputil.NewRequest(ctx).
		SetHeader("Authorization", ir.AuthValue).
		SetHeader(httputil.RequestIDKey, httputil.GenRequestID(ctx))
}
