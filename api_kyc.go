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

// GetKycProfileByUiamID GetKycStatusByUiamID
func (ir IDRequest) GetKycProfileByUiamID(ctx context.Context, uiamID uint64) (*Profile, error) {
	var profile = new(Profile)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s%s?id=%v", ir.ServerURL, "/v1/kyc/profile", uiamID), nil, profile); err != nil {
		return nil, err
	}

	profile.KycErrorMessage = FaceidKycResult[profile.KycError]

	return profile, nil
}
