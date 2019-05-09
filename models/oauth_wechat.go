package models

// OauthWechat 微信登陆
type OauthWechat struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

// TableName 设置表名
func (model *OauthWechat) TableName() string {
	return "oauth_wechat"
}
