package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfEventLog
//@description: 创建AvfEventLog记录
//@param: avfEventLog model.AvfEventLog
//@return: err error

func CreateAvfEventLog(avfEventLog model.AvfEventLog) (err error) {
	err = global.GVA_DB.Create(&avfEventLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfEventLog
//@description: 删除AvfEventLog记录
//@param: avfEventLog model.AvfEventLog
//@return: err error

func DeleteAvfEventLog(avfEventLog model.AvfEventLog) (err error) {
	err = global.GVA_DB.Delete(&avfEventLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfEventLogByIds
//@description: 批量删除AvfEventLog记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfEventLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfEventLog{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfEventLog
//@description: 更新AvfEventLog记录
//@param: avfEventLog *model.AvfEventLog
//@return: err error

func UpdateAvfEventLog(avfEventLog model.AvfEventLog) (err error) {
	err = global.GVA_DB.Save(&avfEventLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfEventLog
//@description: 根据id获取AvfEventLog记录
//@param: id uint
//@return: err error, avfEventLog model.AvfEventLog

func GetAvfEventLog(id uint) (err error, avfEventLog model.AvfEventLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfEventLog).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfEventLogInfoList
//@description: 分页获取AvfEventLog记录
//@param: info request.AvfEventLogSearch
//@return: err error, list interface{}, total int64

func GetAvfEventLogInfoList(info request.AvfEventLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AvfEventLog{})
    var avfEventLogs []model.AvfEventLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.BlockNumber != "" {
        db = db.Where("`block_number` = ?",info.BlockNumber)
    }
    if info.Contract != "" {
        db = db.Where("`contract` = ?",info.Contract)
    }
    if info.Form != "" {
        db = db.Where("`form` = ?",info.Form)
    }
    if info.To != "" {
        db = db.Where("`to` = ?",info.To)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfEventLogs).Error
	return err, avfEventLogs, total
}