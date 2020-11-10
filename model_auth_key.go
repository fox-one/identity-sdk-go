package identity

import "time"

// AuthorizedKey AuthorizedKey
type AuthorizedKey struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	AppID     uint32         `json:"app_id"`
	UserID    uint64         `json:"user_id"`
	Type      AuthTypeEnum   `json:"type"`
	Scheme    AuthSchemeEnum `json:"scheme"`
	Scopes    string         `json:"scopes"`
	Key       string         `json:"auth_key"`
	Secret    string         `json:"-"`
	Remark    string         `json:"remark"`
	ExpriedAt *time.Time     `json:"expried_at"`
	CreatedAt time.Time      `json:"created_at"`
}







