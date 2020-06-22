package identity

import (
	"time"
)

// AuthorizationResponse AuthorizationResponse
type AuthorizationResponse struct {
	ID         uint64               `json:"id"`
	UserID     uint64               `json:"user_id"`
	AppID      uint32               `json:"app_id"`
	Provider   AuthProviderTypeEnum `json:"provider"`
	OauthID    string               `json:"oauth_id"`
	UnionID    string               `json:"union_id"`
	AppUserID  string               `json:"app_user_id"`
	Credential string               `json:"credential"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}

type AuthorizationListResponse struct {
	Data       []*AuthorizationResponse `json:"data"`
	Pagination Pagination               `json:"pagination"`
}

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
	MixinAuth *MixinAuthResponse `json:"mixin,omitempty"`
	FoxAuth   *FoxAuthResponse   `json:"foxone,omitempty"`
}

// MixinAuthResponse MixinAuthResponse
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

// TokenResponse profile
type TokenResponse struct {
	AppID     uint32         `json:"app_id,omitempty"`
	UserID    uint64         `json:"user_id,omitempty"`
	Type      AuthTypeEnum   `json:"type,omitempty"`
	Scheme    AuthSchemeEnum `json:"scheme,omitempty"`
	Provider  string         `json:"oap,omitempty"`
	Key       string         `json:"sid,omitempty"`
	ExpriedAt *time.Time     `json:"expried_at,omitempty"`
	Token     string         `json:"token,omitempty"`
}

// AuthProviderTypeEnum 枚举
type AuthProviderTypeEnum string

const (
	// AuthProviderTypeEnumMixin offer
	AuthProviderTypeEnumMixin AuthProviderTypeEnum = "mixin"
	// AuthProviderTypeEnumFoxone order
	AuthProviderTypeEnumFoxone AuthProviderTypeEnum = "foxone"
	// AuthProviderTypeEnumWechat wechat
	AuthProviderTypeEnumWechat AuthProviderTypeEnum = "wechat"
	// AuthProviderTypeEnumAlipay alipay
	AuthProviderTypeEnumAlipay AuthProviderTypeEnum = "alipay"
	// AuthProviderTypeEnumUnkonwn other
	AuthProviderTypeEnumUnkonwn AuthProviderTypeEnum = "unkonwn"
)

func (e AuthProviderTypeEnum) String() string {
	switch e {
	case AuthProviderTypeEnumMixin:
		return "mixin"
	case AuthProviderTypeEnumFoxone:
		return "foxone"
	case AuthProviderTypeEnumWechat:
		return "wechat"
	case AuthProviderTypeEnumAlipay:
		return "alipay"
	default:
		return "unkonwn"
	}
}

// AuthTypeEnum 枚举
type AuthTypeEnum string

const (
	// AuthTypeEnumUser user
	AuthTypeEnumUser AuthTypeEnum = "user"
	// AuthTypeEnumApp app
	AuthTypeEnumApp AuthTypeEnum = "app"
	// AuthTypeEnumAppuser appuser
	AuthTypeEnumAppuser AuthTypeEnum = "appuser"
	// AuthTypeEnumSystem system
	AuthTypeEnumSystem AuthTypeEnum = "system"
)

func (e AuthTypeEnum) String() string {
	switch e {
	case AuthTypeEnumUser:
		return "user"
	case AuthTypeEnumApp:
		return "app"
	case AuthTypeEnumAppuser:
		return "appuser"
	case AuthTypeEnumSystem:
		return "system"
	default:
		return "unkonwn"
	}
}

// AuthSchemeEnum 枚举
type AuthSchemeEnum string

const (
	// AuthSchemeEnumBasic basic
	AuthSchemeEnumBasic AuthSchemeEnum = "basic"
	// AuthSchemeEnumDigest digest
	AuthSchemeEnumDigest AuthSchemeEnum = "digest"
	// AuthSchemeEnumJWTHS jwt_hs
	AuthSchemeEnumJWTHS AuthSchemeEnum = "jwths"
	// AuthSchemeEnumJWTRS jwt_rs
	AuthSchemeEnumJWTRS AuthSchemeEnum = "jwtrs"
	// AuthSchemeEnumJWTES jwt_es
	AuthSchemeEnumJWTES AuthSchemeEnum = "jwtes"
)

func (e AuthSchemeEnum) String() string {
	switch e {
	case AuthSchemeEnumBasic:
		return "basic"
	case AuthSchemeEnumDigest:
		return "digest"
	case AuthSchemeEnumJWTHS:
		return "jwths"
	case AuthSchemeEnumJWTRS:
		return "jwtrs"
	case AuthSchemeEnumJWTES:
		return "jwtes"
	default:
		return "unkonwn"
	}
}
