package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfCardTransfer
//@description: 创建AvfCardTransfer记录
//@param: avfCardTransfer model.AvfCardTransfer
//@return: err error

func CreateAvfCardTransfer(avfCardTransfer model.AvfCardTransfer) (err error) {
	err = global.GVA_DB.Create(&avfCardTransfer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfCardTransfer
//@description: 删除AvfCardTransfer记录
//@param: avfCardTransfer model.AvfCardTransfer
//@return: err error

func DeleteAvfCardTransfer(avfCardTransfer model.AvfCardTransfer) (err error) {
	err = global.GVA_DB.Delete(&avfCardTransfer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfCardTransferByIds
//@description: 批量删除AvfCardTransfer记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfCardTransferByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfCardTransfer{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfCardTransfer
//@description: 更新AvfCardTransfer记录
//@param: avfCardTransfer *model.AvfCardTransfer
//@return: err error

func UpdateAvfCardTransfer(avfCardTransfer model.AvfCardTransfer) (err error) {
	err = global.GVA_DB.Save(&avfCardTransfer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfCardTransfer
//@description: 根据id获取AvfCardTransfer记录
//@param: id uint
//@return: err error, avfCardTransfer model.AvfCardTransfer

func GetAvfCardTransfer(id uint) (err error, avfCardTransfer model.AvfCardTransfer) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfCardTransfer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfCardTransferInfoList
//@description: 分页获取AvfCardTransfer记录
//@param: info request.AvfCardTransferSearch
//@return: err error, list interface{}, total int64

func GetAvfCardTransferInfoList(info request.AvfCardTransferSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AvfCardTransfer{})
	var avfCardTransfers []model.AvfCardTransfer
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.From != "" {
		db = db.Where("`from` = ?", info.From)
	}
	if info.To != "" {
		db = db.Where("`to` = ?", info.To)
	}
	if info.Block != "" {
		db = db.Where("`block` = ?", info.Block)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfCardTransfers).Error
	return err, avfCardTransfers, total
}
