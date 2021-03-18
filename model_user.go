package identity

import "time"

// User User
type User struct {
	UserID      uint64 `mapstructure:"id" json:"id"`
	UserName    string `mapstructure:"user_name" json:"user_name"`
	PhoneCode   string `mapstructure:"phone_code" json:"phone_code"`
	PhoneNumber string `mapstructure:"phone_number" json:"phone_number"`
	//MfaCredential *MfaCredential `mapstructure:"mfa_credential" json:"mfa_credential"`
	Email       string      `mapstructure:"email" json:"email"`
	Description string      `mapstructure:"description" json:"description"`
	Attributes  string      `mapstructure:"attributes" json:"attributes"`
	Auths       *AuthObject `mapstructure:"authorizations" json:"authorizations,omitempty"`
	Profile     *Profile    `mapstructure:"profile" json:"profile,omitempty"`
	Status      string      `mapstructure:"status" json:"status"`
	CreatedAt   time.Time   `mapstructure:"created_at" json:"created_at"`
}

// Profile profile
type Profile struct {
	UserID        uint64     `json:"user_id"`
	Name          string     `json:"name,omitempty"`
	KycLevel      int        `json:"kyc_level,omitempty"`
	KycStatus     string     `json:"kyc_status,omitempty"`
	KycError      string     `json:"kyc_error,omitempty"`
	IDDigest      string     `json:"id_digest,omitempty"`
	IDType        IDTypeEnum `json:"id_type,omitempty"`
	IDNoEncrypted string     `json:"id_no_encrypted,omitempty"`
	CertifiedAt   time.Time  `json:"certified_at"`
}

// MfaCredentialResponse MfaCredentialResponse
type MfaCredential struct {
	Type            MfaTypeEnum `json:"type"`
	Enable          bool        `json:"enable"`
	ProvisioningUri string      `json:"provisioning_uri,omitempty"`
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

// =============== IDTypeEnum =============== //

// IDTypeEnum 证件类型 枚举
type IDTypeEnum string

const (
	// IDTypeEnumIDCard 身份证
	IDTypeEnumIDCard IDTypeEnum = "idcard"
	// IDTypeEnumDriverLicense 驾照
	IDTypeEnumDriverLicense IDTypeEnum = "driverlicense"
	// IDTypeEnumPassport 护照
	IDTypeEnumPassport IDTypeEnum = "passport"
	// IDTypeEnumPermanentResident 户口本
	IDTypeEnumPermanentResident IDTypeEnum = "permanentresident"
	// IDTypeEnumForeign 外国人永久居留证
	IDTypeEnumForeign IDTypeEnum = "foreign"
	// IDTypeEnumArmymanCard 军人证
	IDTypeEnumArmymanCard IDTypeEnum = "armymancard"
	// IDTypeEnumPoliceCard 武警证
	IDTypeEnumPoliceCard IDTypeEnum = "policecard"
	// IDTypeEnumCachet 公章
	IDTypeEnumCachet IDTypeEnum = "cachet"
	// IDTypeEnumBusinessLicense 工商营业执照
	IDTypeEnumBusinessLicense IDTypeEnum = "businesslicense"
	// IDTypeEnumCorporationID 法人代码证
	IDTypeEnumCorporationID IDTypeEnum = "corporationid"
	// IDTypeEnumStudentCard 学生证
	IDTypeEnumStudentCard IDTypeEnum = "studentcard"
	// IDTypeEnumSoldierCard 士兵证
	IDTypeEnumSoldierCard IDTypeEnum = "soldiercard"
	// IDTypeEnumGAJMLW 港澳居民来往内地通行证
	IDTypeEnumGAJMLW IDTypeEnum = "gajmlw"
	// IDTypeEnumTWJMLW 台湾居民来往大陆通行证
	IDTypeEnumTWJMLW IDTypeEnum = "twjmlw"
)
