package identity

import "time"

// User User
type User struct {
	ID          uint64    `json:"id"`
	APPID       uint64    `json:"app_id"`
	PhoneCode   string    `json:"phone_code"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserName    string    `json:"user_name"`
	AvatarURL   string    `json:"avatar_url"`
	Description string    `json:"description"`
	Remark      string    `json:"remark"`
	Mixin       *Provider `json:"mixin"`
	FoxOne      *Provider `json:"foxone"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Provider Provider
type Provider struct {
	Provider   string    `json:"provider"`
	OauthID    string    `json:"oauth_id"`
	UnionID    string    `json:"union_id"`
	MixinID    string    `json:"mixin_id"`
	FoxOneID   string    `json:"foxone_id"`
	Credential string    `json:"credential"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Profile Profile
type Profile struct {
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name"`
	KycLevel    int       `json:"kyc_level"`
	KycStatus   string    `json:"kyc_status"`
	KycError    string    `json:"kyc_error"`
	CertifiedAt time.Time `json:"certified_at"`
	IDDigest    string    `json:"id_digest"`
}
