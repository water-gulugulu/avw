package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfCard
//@description: 创建AvfCard记录
//@param: avfCard model.AvfCard
//@return: err error

func CreateAvfCard(avfCard model.AvfCard) (err error) {
	err = global.GVA_DB.Create(&avfCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfCard
//@description: 删除AvfCard记录
//@param: avfCard model.AvfCard
//@return: err error

func DeleteAvfCard(avfCard model.AvfCard) (err error) {
	err = global.GVA_DB.Delete(&avfCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfCardByIds
//@description: 批量删除AvfCard记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfCardByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfCard{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfCard
//@description: 更新AvfCard记录
//@param: avfCard *model.AvfCard
//@return: err error

func UpdateAvfCard(avfCard model.AvfCard) (err error) {
	err = global.GVA_DB.Save(&avfCard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfCard
//@description: 根据id获取AvfCard记录
//@param: id uint
//@return: err error, avfCard model.AvfCard

func GetAvfCard(id uint) (err error, avfCard model.AvfCard) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfCard).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfCardInfoList
//@description: 分页获取AvfCard记录
//@param: info request.AvfCardSearch
//@return: err error, list interface{}, total int64

func GetAvfCardInfoList(info request.AvfCardSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AvfCard{})
    var avfCards []model.AvfCard
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfCards).Error
	return err, avfCards, total
}