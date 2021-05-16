package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfOrderCard
//@description: 创建AvfOrderCard记录
//@param: avfOrderCard model.AvfOrderCard
//@return: err error

func CreateAvfOrderCard(avfOrderCard model.AvfOrderCard) (err error) {
	err = global.GVA_DB.Create(&avfOrderCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfOrderCard
//@description: 删除AvfOrderCard记录
//@param: avfOrderCard model.AvfOrderCard
//@return: err error

func DeleteAvfOrderCard(avfOrderCard model.AvfOrderCard) (err error) {
	err = global.GVA_DB.Delete(&avfOrderCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfOrderCardByIds
//@description: 批量删除AvfOrderCard记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfOrderCardByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfOrderCard{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfOrderCard
//@description: 更新AvfOrderCard记录
//@param: avfOrderCard *model.AvfOrderCard
//@return: err error

func UpdateAvfOrderCard(avfOrderCard model.AvfOrderCard) (err error) {
	err = global.GVA_DB.Save(&avfOrderCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfOrderCard
//@description: 根据id获取AvfOrderCard记录
//@param: id uint
//@return: err error, avfOrderCard model.AvfOrderCard

func GetAvfOrderCard(id uint) (err error, avfOrderCard model.AvfOrderCard) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfOrderCard).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfOrderCardInfoList
//@description: 分页获取AvfOrderCard记录
//@param: info request.AvfOrderCardSearch
//@return: err error, list interface{}, total int64

func GetAvfOrderCardInfoList(info request.AvfOrderCardSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AvfOrderCard{})
	var avfOrderCards []model.AvfOrderCard
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != 0 {
		db = db.Where("`order_id` = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfOrderCards).Error
	return err, avfOrderCards, total
}
