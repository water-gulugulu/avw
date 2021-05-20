// 自动生成模板AvfUser
package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfUser struct {
	global.GVA_MODEL
	Pid           string `json:"pid" form:"pid" gorm:"column:pid;comment:用户的上级地址;type:varchar(255);size:255;"`                             // 上级地址
	Username      string `json:"username" form:"username" gorm:"column:username;comment:用户名;type:varchar(20);size:20;"`                    // 用户名
	Mobile        string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:帐号手机号;type:char;"`                                       // 手机号
	WalletAddress string `json:"walletAddress" form:"walletAddress" gorm:"column:wallet_address;comment:钱包地址;type:varchar(100);size:100;"` // 钱包地址
	Password      string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar(32);size:32;"`                     // 密码
	PayPassword   string `json:"payPassword" form:"payPassword" gorm:"column:pay_password;comment:支付密码;type:varchar(32);size:32;"`         // 支付密码
	LoginTime     int    `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:登录时间;type:int;size:10;"`                       // 登录时间
	LoginIp       string `json:"loginIp" form:"loginIp" gorm:"column:login_ip;comment:登录ip;type:varchar(30);size:30;"`                     // 登录IP
	LoginTimes    int    `json:"loginTimes" form:"loginTimes" gorm:"column:login_times;comment:登录次数;type:int;size:10;"`                    // 登录次数
	CreatedTime   int    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:创建时间;type:int;size:10;"`                 // 创建时间
	Status        *bool  `json:"status" form:"status" gorm:"column:status;comment:状态 1-正常 0-禁用;type:tinyint;"`                             // 状态
}

func (h *AvfUser) TableName() string {
	return "avf_user"
}

func (h *AvfUser) FindUserByAddress(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("wallet_address = ?", h.WalletAddress).First(&h).Error
}

func (h *AvfUser) FindUserID(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id = ?", h.ID).First(&h).Error
}
func (h *AvfUser) CreateUser(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}
func (h *AvfUser) FindUserByPid(DB *gorm.DB) (list []AvfUser, err error) {
	if err = DB.Table(h.TableName()).Where("pid = ?", h.WalletAddress).Select("id,pid,username,wallet_address,created_at").Find(&list).Error; err != nil {
		return nil, err
	}

	return
}

func (h *AvfUser) GetListAll(DB *gorm.DB) (list []*AvfUser, err error) {

	if err = DB.Table(h.TableName()).Order("id asc").Find(&list).Error; err != nil {
		return nil, err
	}
	return
}
