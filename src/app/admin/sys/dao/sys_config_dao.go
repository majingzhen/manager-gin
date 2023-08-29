// Package model 自动生成模板 SysConfigDao
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_config/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysConfigDao 结构体

type SysConfigDao struct{}

// Create 创建SysConfig记录
// Author
func (dao *SysConfigDao) Create(sysConfig model.SysConfig) (err error) {
	err = global.GOrmDao.Create(&sysConfig).Error
	return err
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (dao *SysConfigDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysConfig{}, "id in ?", ids).Error
	return err
}

// Update 更新SysConfig记录
// Author
func (dao *SysConfigDao) Update(sysConfig model.SysConfig) (err error) {
	err = global.GOrmDao.Updates(&sysConfig).Error
	return err
}

// Get 根据id获取SysConfig记录
// Author
func (dao *SysConfigDao) Get(id string) (err error, sysConfig *model.SysConfig) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysConfig).Error
	return
}

// Page 分页获取SysConfig记录
// Author
func (dao *SysConfigDao) Page(param *view.SysConfigPageView) (err error, res *common.PageInfo) {
	// 创建db
	db := global.GOrmDao.Model(&model.SysConfig{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.ConfigName != "" {
		db.Where("config_name like ?", "%"+param.ConfigName+"%")
	}
	if param.ConfigKey != "" {
		db.Where("config_key like ?", "%"+param.ConfigKey+"%")
	}
	if param.ConfigType != "" {
		db.Where("config_type = ?", param.ConfigType)
	}
	page := common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}

	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.SysConfig
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取SysConfig记录
// Author
func (dao *SysConfigDao) List(data *model.SysConfig) (err error, datas []*model.SysConfig) {
	var rows []*model.SysConfig
	db := global.GOrmDao.Model(&model.SysConfig{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = rows
	return err, datas
}

// SelectConfigByKey 根据key查询SysConfig记录
func (dao *SysConfigDao) SelectConfigByKey(key string) (error, *model.SysConfig) {
	var rows []*model.SysConfig
	db := global.GOrmDao.Model(&model.SysConfig{})
	db.Where("config_key = ?", key)
	err := db.Find(&rows).Error
	if rows != nil && len(rows) > 0 {
		return err, rows[0]
	}
	return err, nil
}
