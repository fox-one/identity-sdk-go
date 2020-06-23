package identity

import "time"

// User User
type User struct {
	UserID      uint64      `json:"id"`
	UserName    string      `json:"user_name"`
	PhoneCode   string      `json:"phone_code"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Description string      `json:"description"`
	Auths       *AuthObject `json:"authorizations,omitempty"`
	Profile     *Profile    `json:"profile,omitempty"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}

// Profile profile
type Profile struct {
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name,omitempty"`
	KycLevel    int       `json:"kyc_level,omitempty"`
	KycStatus   string    `json:"kyc_status,omitempty"`
	KycError    string    `json:"kyc_error,omitempty"`
	IDDigest    string    `json:"id_digest,omitempty"`
	CertifiedAt time.Time `json:"certified_at"`
}
