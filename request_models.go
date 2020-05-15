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

type KycSyncRequest struct {
	UserID      uint64    `json:"user_id"`
	KycLevel    int       `json:"kyc_level"`
	KycStatus   string    `json:"kyc_status"`
	KycError    string    `json:"kyc_error"`
	CertifiedAt time.Time `json:"certified_at"`
	IDNo        string    `json:"id_no"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	IDType      string    `json:"id_type"`
}
