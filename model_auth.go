package identity

import (
	"time"
)

// Authorization Authorization
type Authorization struct {
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

// AuthorizationList AuthorizationList
type AuthorizationList struct {
	Data       []*Authorization `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

// AuthObject Auths
type AuthObject struct {
	MixinAuth *MixinAuth `json:"mixin,omitempty"`
	FoxAuth   *FoxAuth   `json:"foxone,omitempty"`
}

// MixinAuth MixinAuth
type MixinAuth struct {
	UserID     uint64      `json:"user_id"`
	Provider   string      `json:"provider"`
	OauthID    string      `json:"oauth_id"`
	MixinID    string      `json:"mixin_id"`
	Credential interface{} `json:"credential"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// FoxAuth FoxAuth
type FoxAuth struct {
	UserID     uint64            `json:"user_id"`
	Provider   string            `json:"provider"`
	OauthID    string            `json:"oauth_id"`
	FoxID      string            `json:"foxone_id"`
	Credential *FoxoneCredential `json:"credential"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// FoxoneCredential FoxoneCredential
type FoxoneCredential struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// MixinTokenCredential MixinCredential
type MixinTokenCredential struct {
	Type        MixinCredentialTypeEnum `json:"type"`
	AccessToken string                  `json:"access_token"`
	Scope       string                  `json:"scope"`
}

// MixinEdkeyCredential MixinCredential
type MixinEdkeyCredential struct {
	Type           MixinCredentialTypeEnum `json:"type"`
	EdPrivKey      string                  `json:"ed_priv_key"`
	EdServerPubKey string                  `json:"ed_server_pub_key"`
	ClientID       string                  `json:"client_id"`
	AuthID         string                  `json:"auth_id"`
	Scope          string                  `json:"scope"`
}

// ================ 枚举 ================== //
// ================ 枚举 ================== //

// MixinCredentialTypeEnum 枚举
type MixinCredentialTypeEnum string

const (
	// MixinCredentialTypeEnumEdkey edkey
	MixinCredentialTypeEnumEdkey MixinCredentialTypeEnum = "edkey"
	// MixinCredentialTypeEnumToken token
	MixinCredentialTypeEnumToken MixinCredentialTypeEnum = "token"
)

func (e MixinCredentialTypeEnum) String() string {
	switch e {
	case MixinCredentialTypeEnumEdkey:
		return "edkey"
	case MixinCredentialTypeEnumToken:
		return "token"
	default:
		return ""
	}
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
