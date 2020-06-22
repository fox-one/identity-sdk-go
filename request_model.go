package identity

import "time"

// CreateUserReq CreateUserReq
type CreateUserReq struct {
	UserName    string `json:"username"`
	Type        string `json:"type"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AvatarURL   string `json:"avatar_url"`
	Description string `json:"description"`
}

// MixinAuthReq MixinAuthReq
type MixinAuthReq struct {
	UserName    string `json:"username"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	MixinID     string `json:"mixin_id"`
	OauthID     string `json:"oauth_id"`
	Credential  string `json:"credential"`
	AvatarURL   string `json:"avatar_url"`
}

// FoxoneAuthReq FoxoneAuthReq
type FoxoneAuthReq struct {
	UserName    string `json:"username"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	FoxoneID    string `json:"foxone_id"`
	OauthID     string `json:"oauth_id"`
	Credential  string `json:"credential"`
	AvatarURL   string `json:"avatar_url"`
}

// LoginRequest LoginRequest
type LoginRequest struct {
	ID          uint64 `json:"id"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// KycSyncRequest KycSyncRequest
type KycSyncRequest struct {
	UserID      uint64     `json:"user_id"`
	KycLevel    int        `json:"kyc_level"`
	KycStatus   string     `json:"kyc_status"`
	KycError    string     `json:"kyc_error"`
	CertifiedAt *time.Time `json:"certified_at"`
	IDNo        string     `json:"id_no"`
	Name        string     `json:"name"`
	Country     string     `json:"country"`
	IDType      string     `json:"id_type"`
}

// KycBizTokenRequest KycBizTokenRequest
type KycBizTokenRequest struct {
	UserID    string `json:"user_id"`
	ReturnURL string `json:"return_url"`
}

// AuthListRequest 用户表求
type AuthListRequest struct {
	Offset   uint64               `json:"offset" `
	Limit    uint64               `json:"limit"`
	Provider AuthProviderTypeEnum `json:"provider"`
}

type TokenCreateRequest struct {
	UserID    uint64         `json:"user_id"`
	Type      AuthTypeEnum   `json:"type"`
	Scheme    AuthSchemeEnum `json:"scheme"`
	Provider  string         `json:"oap"`
	SessionID string         `json:"sid"`
	Sign      string         `json:"sig"`
	SignAlg   string         `json:"sal"`
	Duration  time.Duration  `json:"duration"` //  有效时长
}