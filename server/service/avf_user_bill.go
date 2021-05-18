package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfUserBill
//@description: 创建AvfUserBill记录
//@param: avfUserBill model.AvfUserBill
//@return: err error

func CreateAvfUserBill(avfUserBill model.AvfUserBill) (err error) {
	err = global.GVA_DB.Create(&avfUserBill).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfUserBill
//@description: 删除AvfUserBill记录
//@param: avfUserBill model.AvfUserBill
//@return: err error

func DeleteAvfUserBill(avfUserBill model.AvfUserBill) (err error) {
	err = global.GVA_DB.Delete(&avfUserBill).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfUserBillByIds
//@description: 批量删除AvfUserBill记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfUserBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfUserBill{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfUserBill
//@description: 更新AvfUserBill记录
//@param: avfUserBill *model.AvfUserBill
//@return: err error

func UpdateAvfUserBill(avfUserBill model.AvfUserBill) (err error) {
	err = global.GVA_DB.Save(&avfUserBill).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfUserBill
//@description: 根据id获取AvfUserBill记录
//@param: id uint
//@return: err error, avfUserBill model.AvfUserBill

func GetAvfUserBill(id uint) (err error, avfUserBill model.AvfUserBill) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfUserBill).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfUserBillInfoList
//@description: 分页获取AvfUserBill记录
//@param: info request.AvfUserBillSearch
//@return: err error, list interface{}, total int64

func GetAvfUserBillInfoList(info request.AvfUserBillSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AvfUserBill{})
	var avfUserBills []model.AvfUserBill
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfUserBills).Error
	return err, avfUserBills, total
}
