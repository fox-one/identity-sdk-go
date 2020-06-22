package identity

import (
	"context"
	"fmt"

	httputil "github.com/fox-one/identity-sdk-go/utils"
)

// GetHHBizToken GetHHBizToken
func (ir IDRequest) GetHHBizToken(ctx context.Context, userID uint64, returnURL string) (string, error) {
	var result map[string]interface{}

	if err := httputil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s%s?user_id=%v&return_url=%s", ir.ServerURL, "/v1/kyc/hh/faceid/token", userID, returnURL), nil, &result); err != nil {
		return "", err
	}

	return result["url"].(string), nil
}

// GetZDBizToken GetZDBizToken
func (ir IDRequest) GetZDBizToken(ctx context.Context, userID uint64, returnURL string) (string, error) {
	var result map[string]interface{}

	if err := httputil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s%s?user_id=%v&return_url=%s", ir.ServerURL, "/v1/kyc/zd/faceid/token", userID, returnURL), nil, &result); err != nil {
		return "", err
	}

	return result["url"].(string), nil
}
