package identity

import "time"

// User User
type User struct {
	UserID        uint64         `mapstructure:"id" json:"id"`
	UserName      string         `mapstructure:"user_name" json:"user_name"`
	PhoneCode     string         `mapstructure:"phone_code" json:"phone_code"`
	PhoneNumber   string         `mapstructure:"phone_number" json:"phone_number"`
	MfaCredential *MfaCredential `mapstructure:"mfa_credential" json:"mfa_credential"`
	Email         string         `mapstructure:"email" json:"email"`
	Description   string         `mapstructure:"description" json:"description"`
	Attributes    string         `mapstructure:"attributes" json:"attributes"`
	Auths         *AuthObject    `mapstructure:"authorizations" json:"authorizations,omitempty"`
	Profile       *Profile       `mapstructure:"profile" json:"profile,omitempty"`
	Status        string         `mapstructure:"status" json:"status"`
	CreatedAt     time.Time      `mapstructure:"created_at" json:"created_at"`
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

// MfaCredentialResponse MfaCredentialResponse
type MfaCredential struct {
	Type   MfaTypeEnum `json:"type"`
	Enable bool        `json:"enable"`
}

type MfaTypeEnum string

const (
	MfaTypeEnumSms   = "sms"
	MfaTypeEnumEmail = "email"
	MfaTypeEnumTotp  = "totp"
	MfaTypeEnumHotp  = "hotp"
)

func (e MfaTypeEnum) String() string {
	switch e {
	case MfaTypeEnumSms:
		return "sms"
	case MfaTypeEnumEmail:
		return "email"
	case MfaTypeEnumTotp:
		return "totp"
	case MfaTypeEnumHotp:
		return "hotp"
	default:
		return "unkonwn"
	}
}
