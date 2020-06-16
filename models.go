package identity

import "time"

// User User
type User struct {
	ID          uint64    `json:"id"`
	APPID       uint64    `json:"app_id"`
	PhoneCode   string    `json:"phone_code"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserName    string    `json:"user_name"`
	AvatarURL   string    `json:"avatar_url"`
	Description string    `json:"description"`
	Remark      string    `json:"remark"`
	Mixin       *Provider `json:"mixin"`
	FoxOne      *Provider `json:"foxone"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Provider Provider
type Provider struct {
	Provider   string    `json:"provider"`
	OauthID    string    `json:"oauth_id"`
	UnionID    string    `json:"union_id"`
	MixinID    string    `json:"mixin_id"`
	FoxOneID   string    `json:"foxone_id"`
	Credential string    `json:"credential"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Profile Profile
type Profile struct {
	UserID          uint64    `json:"user_id"`
	Name            string    `json:"name"`
	KycLevel        int       `json:"kyc_level"`
	KycStatus       string    `json:"kyc_status"`
	KycError        string    `json:"kyc_error"`
	KycErrorMessage string    `json:"kyc_error_message"`
	CertifiedAt     time.Time `json:"certified_at"`
	IDDigest        string    `json:"id_digest"`
}

// FaceidKycResult FaceidKycResult
var FaceidKycResult = map[string]string{
	// 活体成功
	"1000:SUCCESS":         "待比对照片与参考数据照片或上传照片对比是同一个人",
	"1000:LIVENESS_FINISH": "活体完成",

	// 身份证相关 - 正确
	"1001:SUCCESS":                  "尚未进行活体认证", // 识别出身份证人像面和国徽面均没有问题"
	"1001:SUCCESS_BACKSIDE_SINGLE":  "只需识别国徽面时识别成功",
	"1001:SUCCESS_FRONTSIDE_SINGLE": "只需识别人像面时识别成功",

	// 身份证相关 - 错误
	"1002:SUCCESS_ALMOST_FRONTSIDE":        "识别出身份证国徽面没有问题，身份证人像面内容上存在没有识别出来、或者出来的内容存在逻辑问题、识别出来的内容存在质量问题",
	"1002:SUCCESS_ALMOST_BACKSIDE":         "识别出身份证人像面没有问题，身份证国徽面内容上存在没有识别出来、或者出来的内容存在逻辑问题、识别出来的内容存在质量问",
	"1002:SUCCESS_ALMOST_BOTHSIDE":         "识别出身份证人像面和国徽面内容上均存在没有识别出来、或者出来的内容存在逻辑问题、识别出来的内容存在质量问题",
	"1002:SUCCESS_ALMOST_BACKSIDE_SINGLE":  "只需识别国徽面时内容上存在没有识别出来、或者出来的内容存在逻辑问题、识别出来的内容存在质量问题",
	"1002:SUCCESS_ALMOST_FRONTSIDE_SINGLE": "只需识别人像面时内容上存在没有识别出来、或者出来的内容存在逻辑问题、识别出来的内容存在质量问题",
	"2001:NO_IDCARD_FRONTSIDE":             "没有检测到身份证人像面照片",
	"2001:NO_IDCARD_BACKSIDE":              "没有检测到身份证国徽面照片",
	"2001:NO_IDCARD_BOTHSIDE":              "没有检测到身份证人像面和国徽面照片",
	"2001:NO_IDCARD_BACKSIDE_SINGLE":       "只需识别国徽面时没有检测到国徽面照片",
	"2001:NO_IDCARD_FRONTSIDE_SINGLE":      "只需识别人像面时没有检测到人像面照片",
	"2000:PASS_LIVING_NOT_THE_SAME":        "待比对照片与参考数据照片或上传照片对比不是同一个人",
	"3000:NO_ID_CARD_NUMBER":               "无此身份证号",
	"3000:ID_NUMBER_NAME_NOT_MATCH":        "身份证号和姓名不相符",
	"3000:NO_ID_PHOTO":                     "无法获取参考照片",
	"3000:NO_FACE_FOUND":                   "参考照片中找不到人脸",
	"3000:PHOTO_FORMAT_ERROR":              "参考照片格式错误",
	"3000:DATA_SOURCE_ERROR":               "参考数据错误",
	"3100:IDCARD_PHOTO_FRONTSIDE":          "身份证人像面识别错误（可能是拍摄模糊，阴影，暗光）或者有逻辑错误（例如性别和身份证号不匹配）",
	"3100:IDCARD_BACKSIDE_BLURRED":         "身份证国徽面识别错误（可能是拍摄模糊，阴影，暗光）或者有逻辑错误（例如性别和身份证号不匹配）",
	"3100:NO_FACE_FOUND_IDCARD":            "身份证人像面找不到人脸",
	"3100:IDCARD_PHOTO_NOTFRONTSIDE":       "非身份证人像面",
	"3100:IDCARD_PHOTO_NOTBACKSIDE":        "非身份证国徽面",
	"3100:IDCARD_PHOTO_INVALID_SIZE":       "身份证图片超过尺寸限制",
	"3200:FAIL_OCR_FAKE_IDCARD":            "假证",

	// 活体 - 错误
	"4100:FAIL_LIVING_FACE_ATTACK":   "云端活体验证失败",
	"4100:CHANGE_FACE_ATTACK":        "活体验证视频中发生了换脸攻击（视频中不是同一个人）",
	"4100:VIDEO_LACK_FRAMES":         "获取到的活体数据故障，请换一台手机重试",
	"4100:FAIL_EYES_CLOSE_DETECTION": "未通过闭眼检测，活体失败",
	"4200:NO_FACE_FOUND":             "活体验证视频中没有检测到人脸",
	"4200:FACE_QUALITY_TOO_LOW":      "活体验证视频中质量太差",
	"4200:INVALID_VIDEO_DURATION":    "活体验证视频中长度不符合要求（2s～20s）",
	"4200:VIDEO_TOO_LARGE":           "活体验证视频过大",
	"4200:SR_ERROR":                  "活体验证视频中，用户读数语音不符合要求",
	"4200:NOT_SYNCHRONIZED":          "活体验证视频中，用户读数唇语不符合要求",
	"4200:NO_AUDIO":                  "活体验证视频无声音",
	"4200:VIDEO_FORMAT_UNSUPPORTED":  "活体验证视频格式无法识别",
	"4200:LIP_VOICE_NOT_SYNC":        "活体验证视频中语音唇语不同步",
	"4200:VIDEO_OK":                  "活体验证视频可用",
	"4200:VIDEO_MANY_TIMES":          "活体验证视频上传超过阈值（默认为3，get_biz_token接口中liveness_retry_count参数可以设置）",
	"4200:VIDEO_INTERNAL_ERROR":      "活体验证内部错误",
	"4200:BIZ_TOKEN_DENIED":          "传入的 biz_token 不符合要求",
	"4200:AUTHENTICATION_FAIL":       "鉴权失败",
	"4200:MOBILE_PHONE_NOT_SUPPORT":  "手机在不支持列表里",
	"4200:SDK_TOO_OLD":               "SDK版本过旧，已经不被支持",
	"4200:MOBILE_PHONE_NO_AUTHORITY": "没有权限（运动传感器、存储、相机）",
	"4200:USER_CANCELLATION":         "用户活体失败，可能原因：用户取消了",
	"4200:USER_TIMEOUT":              "用户活体失败，可能原因：验证过程超时",
	"4200:VERIFICATION_FAILURE":      "用户活体失败，可能原因：验证失败",
	"4200:UNDETECTED_FACE":           "用户活体失败，可能原因：未检测到人脸",
	"4200:ACTION_ERROR":              "用户活体失败，可能原因：用户动作错误；",
	"9000:LIVING_NOT_START":          "活体验证没有开始",
	"9000:LIVING_IN_PROGRESS":        "正在进行验证",
	"9000:LIVING_OVERTIME":           "操作超时，由于用户在长时间没有进行操作",

	// enterprise
	"5000:API_KEY_BE_DISCONTINUED":      "api_key被停用",
	"5000:IP_NOT_ALLOWED":               "不允许访问的IP",
	"5000:NON_ENTERPRISE_CERTIFICATION": "客户未进行企业认证",
	"5000:ACCOUNT_DISCONTINUED":         "用户帐号已停用",
	"5000:EXPIRED_SIGN":                 "签名过期",
	"5000:INVALID_SIGN":                 "无效的签名",
	"5000:REPLAY_ATTACK":                "重放攻击，单次有效的签名被多次使用",
	"5000:BALANCE_NOT_ENOUGH":           "余额不足",
	"5000:MORE_RETRY_TIMES":             "超过重试次数",
	"5000:ACCOUNT_DISABLED":             "账户已停用",

	// User - Error
	"6000:NO_NETWORK_PERMISSION":  "没有网络权限",
	"6000:NO_CAMERA_PERMISSION":   "没有相机权限",
	"6000:DEVICE_NOT_SUPPORT":     "无法启动相机，请确认摄像头功能完好",
	"6000:ILLEGAL_PARAMETER":      "传入参数不合法",
	"6000:INVALID_BUNDLE_ID":      "信息验证失败，请重启程序或设备后重试",
	"6000:USER_CANCEL":            "用户主动退出流程",
	"6000:NETWORK_ERROR":          "连不上互联网，请连接上互联网后重试",
	"6000:FACE_INIT_FAIL":         "无法启动人脸识别，请稍后重试",
	"6000:LIVENESS_DETECT_FAILED": "活体检测不通过",
	"6000:NO_SENSOR_PERMISSION":   "无法读取运动数据的权限，请开启权限后重试",
	"6000:INIT_FAILED":            "初始化失败",
}
