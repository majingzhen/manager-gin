// Package dao 自动生成模板 TableColumnDao
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/gen/model"
	"manager-gin/src/app/admin/gen/service/table_column/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// TableColumnDao 结构体

type TableColumnDao struct{}

// Create 创建TableColumn记录
// Author
func (dao *TableColumnDao) Create(tx *gorm.DB, genTableColumn *model.TableColumn) (err error) {
	err = tx.Create(genTableColumn).Error
	return err
}

// DeleteByIds 批量删除TableColumn记录
// Author
func (dao *TableColumnDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.TableColumn{}, "id in ?", ids).Error
	return err
}

// Update 更新TableColumn记录
// Author
func (dao *TableColumnDao) Update(genTableColumn model.TableColumn) (err error) {
	err = global.GOrmDao.Updates(&genTableColumn).Error
	return err
}

// Get 根据id获取TableColumn记录
// Author
func (dao *TableColumnDao) Get(id string) (err error, genTableColumn *model.TableColumn) {
	err = global.GOrmDao.Where("id = ?", id).First(&genTableColumn).Error
	return
}

// Page 分页获取TableColumn记录
// Author
func (dao *TableColumnDao) Page(param *view.TableColumnPageView) (err error, page *common.PageInfo) {
	db := global.GOrmDao.Model(&model.TableColumn{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if param.Id != "" {
	//	db.Where("ID = ?", param.Id)
	//}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.TableColumn
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取TableColumn记录
// Author
func (dao *TableColumnDao) List(v *view.TableColumnQueryView) (err error, dataList []*model.TableColumn) {
	db := global.GOrmDao.Model(&model.TableColumn{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	// db.Order("create_time desc")
	err = db.Find(&dataList).Error
	return err, dataList
}

func (dao *TableColumnDao) SelectDbTableColumns(tx *gorm.DB, tableName string) (err error, dataList []*model.TableColumn) {
	db := tx.Table("information_schema.columns")
	db.Select("COLUMN_NAME column_name,( CASE WHEN ( IS_NULLABLE = 'no' && COLUMN_KEY != 'PRI' ) THEN '1' ELSE NULL END ) AS is_required,( CASE WHEN COLUMN_KEY = 'PRI' THEN '1' ELSE '0' END ) AS is_pk,ORDINAL_POSITION AS sort,COLUMN_COMMENT column_comment,( CASE WHEN extra = 'AUTO_INCREMENT' THEN '1' ELSE '0' END ) AS is_increment,COLUMN_TYPE column_type,DATA_TYPE data_type,CHARACTER_MAXIMUM_LENGTH character_maximum_length")
	db.Where("table_schema = (SELECT DATABASE())")
	db.Where("table_name = ?", tableName).Find(&dataList)
	err = db.Find(&dataList).Error
	return err, dataList
}

// DeleteByTableIds 根据表id删除列信息
func (dao *TableColumnDao) DeleteByTableIds(tx *gorm.DB, tableIds []string) error {
	return tx.Table("gen_table_column").Where("table_id in (?)", tableIds).Delete(&model.TableColumn{}).Error
}

// GetColumnListByTableId 根据表id获取列信息
func (dao *TableColumnDao) GetColumnListByTableId(tableId string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GOrmDao.Table("gen_table_column").Where("table_id = ?", tableId).Order("sort").Find(&columns).Error
	return err, columns
}
