package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfOrder
//@description: 创建AvfOrder记录
//@param: avfOrder model.AvfOrder
//@return: err error

func CreateAvfOrder(avfOrder model.AvfOrder) (err error) {
	err = global.GVA_DB.Create(&avfOrder).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfOrder
//@description: 删除AvfOrder记录
//@param: avfOrder model.AvfOrder
//@return: err error

func DeleteAvfOrder(avfOrder model.AvfOrder) (err error) {
	err = global.GVA_DB.Delete(&avfOrder).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfOrderByIds
//@description: 批量删除AvfOrder记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfOrder{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfOrder
//@description: 更新AvfOrder记录
//@param: avfOrder *model.AvfOrder
//@return: err error

func UpdateAvfOrder(avfOrder model.AvfOrder) (err error) {
	err = global.GVA_DB.Save(&avfOrder).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfOrder
//@description: 根据id获取AvfOrder记录
//@param: id uint
//@return: err error, avfOrder model.AvfOrder

func GetAvfOrder(id uint) (err error, avfOrder model.AvfOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfOrder).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfOrderInfoList
//@description: 分页获取AvfOrder记录
//@param: info request.AvfOrderSearch
//@return: err error, list interface{}, total int64

func GetAvfOrderInfoList(info request.AvfOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AvfOrder{})
	var avfOrders []model.AvfOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TxHash != "" {
		db = db.Where("`tx_hash` = ?", info.TxHash)
	}
	if info.Block != "" {
		db = db.Where("`block` = ?", info.Block)
	}
	if info.From != "" {
		db = db.Where("`from` = ?", info.From)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfOrders).Error
	return err, avfOrders, total
}
