// 自动生成模板AvfEventLog
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type AvfEventLog struct {
      global.GVA_MODEL
      BlockNumber  string `json:"blockNumber" form:"blockNumber" gorm:"column:block_number;comment:区块编号;type:varchar(40);size:40;"`
      Contract  string `json:"contract" form:"contract" gorm:"column:contract;comment:合约地址;type:varchar(50);size:50;"`
      Form  string `json:"form" form:"form" gorm:"column:form;comment:发起地址;type:varchar(50);size:50;"`
      To  string `json:"to" form:"to" gorm:"column:to;comment:接收地址;type:varchar(50);size:50;"`
      Index  string `json:"index" form:"index" gorm:"column:index;comment:日志索引;type:varchar(50);size:50;"`
      TxIndex  string `json:"txIndex" form:"txIndex" gorm:"column:tx_index;comment:事务索引;type:varchar(50);size:50;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:事件名称;type:varchar(50);size:50;"`
      Number  int `json:"number" form:"number" gorm:"column:number;comment:操作金额;type:bigint;size:19;"`
      Tokens  string `json:"tokens" form:"tokens" gorm:"column:tokens;comment:tokens;type:varchar(500);size:500;"`
      CreateTime  int `json:"createTime" form:"createTime" gorm:"column:create_time;comment:时间戳;type:int;size:10;"`
}


func (AvfEventLog) TableName() string {
  return "avf_event_log"
}

