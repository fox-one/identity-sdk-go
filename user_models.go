package identity

import (
	"time"
)

// UserAuthsResponse User
type UserAuthsResponse struct {
	UserID      uint64           `json:"id"`
	UserName    string           `json:"user_name"`
	PhoneCode   string           `json:"phone_code"`
	PhoneNumber string           `json:"phone_number"`
	Email       string           `json:"email"`
	Description string           `json:"description"`
	Auths       *AuthsResponse   `json:"authorizations,omitempty"`
	Profile     *ProfileResponse `json:"profile,omitempty"`
	Status      string           `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
}

// AuthsResponse AuthsResponse
type AuthsResponse struct {
	MixinAuth interface{} `json:"mixin,omitempty"`
	FoxAuth   interface{} `json:"foxone,omitempty"`
}

type MixinAuthResponse struct {
	UserID     uint64    `json:"user_id"`
	Provider   string    `json:"provider"`
	OauthID    string    `json:"oauth_id"`
	MixinID    string    `json:"mixin_id"`
	Credential string    `json:"credential"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// FoxAuthResponse FoxAuthResponse
type FoxAuthResponse struct {
	UserID     uint64    `json:"user_id"`
	Provider   string    `json:"provider"`
	OauthID    string    `json:"oauth_id"`
	FoxID      string    `json:"foxone_id"`
	Credential string    `json:"credential"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ProfileResponse profile
type ProfileResponse struct {
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name,omitempty"`
	KycLevel    string    `json:"kyc_level,omitempty"`
	KycStatus   string    `json:"kyc_status,omitempty"`
	KycError    string    `json:"kyc_error,omitempty"`
	IDDigest    string    `json:"id_digest,omitempty"`
	CertifiedAt time.Time `json:"certified_at"`
}
