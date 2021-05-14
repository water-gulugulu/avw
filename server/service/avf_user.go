package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAvfUser
//@description: 创建AvfUser记录
//@param: avfUser model.AvfUser
//@return: err error

func CreateAvfUser(avfUser model.AvfUser) (err error) {
	err = global.GVA_DB.Create(&avfUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfUser
//@description: 删除AvfUser记录
//@param: avfUser model.AvfUser
//@return: err error

func DeleteAvfUser(avfUser model.AvfUser) (err error) {
	err = global.GVA_DB.Delete(&avfUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAvfUserByIds
//@description: 批量删除AvfUser记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAvfUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AvfUser{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAvfUser
//@description: 更新AvfUser记录
//@param: avfUser *model.AvfUser
//@return: err error

func UpdateAvfUser(avfUser model.AvfUser) (err error) {
	err = global.GVA_DB.Save(&avfUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfUser
//@description: 根据id获取AvfUser记录
//@param: id uint
//@return: err error, avfUser model.AvfUser

func GetAvfUser(id uint) (err error, avfUser model.AvfUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&avfUser).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAvfUserInfoList
//@description: 分页获取AvfUser记录
//@param: info request.AvfUserSearch
//@return: err error, list interface{}, total int64

func GetAvfUserInfoList(info request.AvfUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AvfUser{})
	var avfUsers []model.AvfUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Pid != "" {
		db = db.Where("`pid` = ?", info.Pid)
	}
	if info.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.WalletAddress != "" {
		db = db.Where("`wallet_address` = ?", info.WalletAddress)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&avfUsers).Error
	return err, avfUsers, total
}

// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: FindUserByID
// @description: 通过id获取用户信息
// @param: id uint
// @return: err error, user *model.AvfUser

func FindUserByID(id uint) (err error, user *model.AvfUser) {
	var u model.AvfUser
	if err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
