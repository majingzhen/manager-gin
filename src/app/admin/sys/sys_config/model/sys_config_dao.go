// Package model 自动生成模板 SysConfigDao
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysConfigDao 结构体

type SysConfigDao struct{}

// Create 创建SysConfig记录
// Author
func (dao *SysConfigDao) Create(sysConfig SysConfig) (err error) {
	err = global.GOrmDao.Create(&sysConfig).Error
	return err
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (dao *SysConfigDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysConfig{}, "id in ?", ids).Error
	return err
}

// Update 更新SysConfig记录
// Author
func (dao *SysConfigDao) Update(sysConfig SysConfig) (err error) {
	err = global.GOrmDao.Updates(&sysConfig).Error
	return err
}

// Get 根据id获取SysConfig记录
// Author
func (dao *SysConfigDao) Get(id string) (err error, sysConfig *SysConfig) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysConfig).Error
	return
}

// Page 分页获取SysConfig记录
// Author
func (dao *SysConfigDao) Page(param *SysConfig, page *common.PageInfo) (err error, datas *[]SysConfig, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysConfig{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.ConfigName != "" {
		model = model.Where("config_name like ?", "%"+param.ConfigName+"%")
	}
	if param.ConfigKey != "" {
		model = model.Where("config_key like ?", "%"+param.ConfigKey+"%")
	}
	if param.ConfigType != "" {
		model = model.Where("config_type = ?", param.ConfigType)
	}
	if err = model.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		model.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysConfig
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysConfig记录
// Author
func (dao *SysConfigDao) List(data *SysConfig) (err error, datas *[]SysConfig) {
	var rows []SysConfig
	db := global.GOrmDao.Model(&SysConfig{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = &rows
	return err, datas
}

// SelectConfigByKey 根据key查询SysConfig记录
func (dao *SysConfigDao) SelectConfigByKey(key string) (error, *SysConfig) {
	var rows SysConfig
	db := global.GOrmDao.Model(&SysConfig{})
	db.Where("config_key = ?", key)
	err := db.First(&rows).Error
	return err, &rows
}
