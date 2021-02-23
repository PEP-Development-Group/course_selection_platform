package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name" gorm:"default:'系统用户'"`

	AuthorityId string `json:"authorityId"`

	Class string `json:"class"`
	PID   string `json:"pid"`

	CancelNums   int `json:"cancel_nums"`
	TotalCredits int `json:"total_credits"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

type AddCancelNums struct {
	Username string `json:"username"`
	Cnt      int    `json:"cnt"`
}
