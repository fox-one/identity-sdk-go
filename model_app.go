package identity

import "time"

// Application Application
type Application struct {
	ID          uint32        `json:"app_id" `
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	CreatorID   uint64        `json:"creator_id"`
	Description string        `json:"description"`
	Status      AppStatusEnum `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// AppStatusEnum 枚举
type AppStatusEnum string

const (
	// AppStatusEnumNormal 正常
	AppStatusEnumNormal AppStatusEnum = "normal"
	// AppStatusEnumDisabled 下架
	AppStatusEnumDisabled AppStatusEnum = "disabled"
)

func (e AppStatusEnum) String() string {
	switch e {
	case AppStatusEnumNormal:
		return "normal"
	case AppStatusEnumDisabled:
		return "disabled"
	default:
		return "unkonwn"
		// return fmt.Sprintf("%f", e)
	}
}
