package identity

// CreateUserReq CreateUserReq
type CreateUserReq struct {
	UserName    string `json:"username"`
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
