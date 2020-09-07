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

// UserModifyReq UserModifyReq
type UserModifyReq struct {
	UserID      uint64 `json:"user_id"`
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
	Email       string `json:"email"`
	Password    string `json:"password"`

	LoginType LoginTypeEnum `json:"login_type"`
	Captcha   string        `json:"captcha"`
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

// TokenCreateRequest TokenCreateRequest
type TokenCreateRequest struct {
	Type   AuthTypeEnum   `json:"type"`
	Scheme AuthSchemeEnum `json:"scheme"`

	Audience  string `json:"aud,omitempty"` // 这个表示 UIAM_ID
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

	// Custom
	Provider  string `json:"oap"` // Mixin / FoxONE
	UID       string `json:"uid"` // 这个 UID 表示的是第三方系统的 ID，像 ZOTC 里的那个自己的
	SessionID string `json:"sid"`
	Sign      string `json:"sig"`
	SignAlg   string `json:"sal"`

	// ExpriedAt
	Duration time.Duration `json:"duration"` //  有效时长
}

// PhoneCodeVerifyRequest PhoneCodeVerifyRequest
type PhoneCodeVerifyRequest struct {
	PhoneCode   string          `json:"phone_code"`
	PhoneNumber string          `json:"phone_number"`
	CaptchaType CaptchaTypeEnum `json:"captcha_type"`
	Code        string          `json:"code"`
}

// AuthBindingRequest  auth 绑定
type AuthBindingRequest struct {
	UserID      uint64               `json:"user_id"`
	Provider    AuthProviderTypeEnum `json:"provider"`
	OauthID     string               `json:"oauth_id" `
	UnionID     string               `json:"union_id" `
	Credential  string               `json:"credential"`
	AppUserName string               `json:"app_user_name"`
	AppUserID   string               `json:"app_user_id"`
}

type LoginTypeEnum string

const (
	LoginTypeEnumPassword LoginTypeEnum = "password"
	LoginTypeEnumPhone    LoginTypeEnum = "phone"
	LoginTypeEnumEmail    LoginTypeEnum = "email"
)

func (e LoginTypeEnum) String() string {
	switch e {
	case LoginTypeEnumPassword:
		return "password"
	case LoginTypeEnumPhone:
		return "phone"
	case LoginTypeEnumEmail:
		return "email"
	default:
		return ""
	}
}

// TwoFactorRequest TwoFactorRequest
type TwoFactorRequest struct {
	UserID   uint64      `json:"user_id"`
	MfaToken string      `json:"mfa_token"`
	MfaType  MfaTypeEnum `json:"mfa_type"`
	Captcha  string      `json:"captcha"`
}

// MfaCredentialRequest MfaCredentialRequest
type MfaCredentialRequest struct {
	UserID      uint64      `json:"user_id"`
	MfaType     MfaTypeEnum `json:"mfa_type"`
	AccountName string      `json:"account_name"`
	IssuerName  string      `json:"issuer_name"`
}

// CaptchaTypeEnum CaptchaTypeEnum
type CaptchaTypeEnum string

const (
	// CaptchaTypeEnumLogin offer
	CaptchaTypeEnumLogin     CaptchaTypeEnum = "login"
	CaptchaTypeEnumTwoFactor CaptchaTypeEnum = "two-factor"
)

func (e CaptchaTypeEnum) String() string {
	switch e {
	case CaptchaTypeEnumLogin:
		return "login"
	default:
		return ""
	}
}
