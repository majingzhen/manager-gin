// Package table 自动生成模板 TableService
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package table

import (
	"errors"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/gen/dao"
	"manager-gin/src/app/admin/gen/model"
	"manager-gin/src/app/admin/gen/service/table/view"
	columm_service "manager-gin/src/app/admin/gen/service/table_column"
	utils2 "manager-gin/src/app/admin/gen/utils"
	"manager-gin/src/common"
	"manager-gin/src/common/constants"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

// Service 结构体
type Service struct {
	tableDao      dao.TableDao
	viewUtils     view.TableViewUtils
	columnDao     dao.TableColumnDao
	columnService columm_service.TableColumnService
}

// Create 创建Table记录
// Author
func (s *Service) Create(tableView *view.TableView) error {
	if err, table := s.viewUtils.View2Data(tableView); err != nil {
		return err
	} else {
		return s.tableDao.Create(global.GormDao, table)
	}
}

// DeleteByIds 批量删除Table记录
// Author
func (s *Service) DeleteByIds(ids []string) (err error) {
	tx := global.GormDao.Begin()
	if err = s.tableDao.DeleteByIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	} else {
		// 删除列信息
		if err := s.columnDao.DeleteByTableIds(tx, ids); err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
			return nil
		}
	}
}

// Update 更新Table记录
// Author
func (s *Service) Update(tableView *view.TableView) error {
	tx := global.GormDao.Begin()
	// 更新
	options := view.TableViewOptions{
		TreeCode:       tableView.TreeCode,
		TreeParentCode: tableView.TreeParentCode,
		TreeName:       tableView.TreeName,
		ParentMenuId:   tableView.ParentMenuId,
		ParentMenuName: tableView.ParentMenuName,
	}
	if jsonBytes, err := json.Marshal(options); err != nil {
		tx.Rollback()
		return err
	} else {
		tableView.Options = string(jsonBytes)
	}
	if err, table := s.viewUtils.View2Data(tableView); err != nil {
		tx.Rollback()
		return err
	} else {
		if err := s.tableDao.Update(tx, table); err != nil {
			tx.Rollback()
			return err
		}
		// 修改列信息
		for _, columnView := range tableView.ColumnList {
			if err := s.columnService.Update(columnView, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	}
}

// Get 根据id获取Table记录
// Author
func (s *Service) Get(id string) (err error, tableView *view.TableView) {
	err1, table := s.tableDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, tableView = s.viewUtils.Data2View(table)
	// options 转字段
	if tableView != nil && tableView.Options != "" {
		var tableOption view.TableViewOptions
		if err := json.Unmarshal([]byte(tableView.Options), &tableOption); err != nil {
			global.Logger.Error("TableOption Convert is error ", zap.Error(err))
		} else {
			tableView.TreeCode = tableOption.TreeCode
			tableView.TreeParentCode = tableOption.TreeParentCode
			tableView.TreeName = tableOption.TreeName
			tableView.ParentMenuId = tableOption.ParentMenuId
			tableView.ParentMenuName = tableOption.ParentMenuName
		}
	}
	return
}

// Page 分页获取Table记录
// Author
func (s *Service) Page(pageInfo *view.TablePageView) (err error, res *common.PageInfo) {
	if err, res = s.tableDao.Page(pageInfo); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// List 获取Table列表
// Author
func (s *Service) List(v *view.TableQueryView) (error, []*view.TableView) {
	if err, dataList := s.tableDao.List(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

// SelectDbTablePage 获取数据库表列表
func (s *Service) SelectDbTablePage(v *view.TablePageView) (err error, res *common.PageInfo) {
	if err, res = s.tableDao.SelectDbTablePage(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// ImportTable 导入Table
func (s *Service) ImportTable(tables string, loginUser string) error {
	tableNames := strings.Split(tables, ",")
	if len(tableNames) == 0 {
		return nil
	}
	if err, tables := s.tableDao.SelectDbTableList(tableNames); err != nil {
		return err
	} else {
		return s.ImportGenTable(tables, loginUser)
	}
}

// ImportGenTable 导入GenTable
func (s *Service) ImportGenTable(tables []*model.Table, loginUser string) error {
	if len(tables) == 0 {
		return nil
	}
	tx := global.GormDao.Begin()
	for _, table := range tables {
		table = utils2.InitTable(table, loginUser)
		table.Id = utils.GenUID()
		if err := s.tableDao.Create(tx, table); err != nil {
			tx.Rollback()
			return err
		}
		// 处理列信息
		if err, tableColumns := s.columnDao.SelectDbTableColumns(tx, table.Name); err != nil {
			tx.Rollback()
			return err
		} else {
			for _, column := range tableColumns {
				column = utils2.InitColumnField(column, table)
				column.Id = utils.GenUID()
				if err := s.columnDao.Create(tx, column); err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	tx.Commit()
	return nil
}

// ValidateEdit 表单验证
func (s *Service) ValidateEdit(v *view.TableView) error {
	if v.TplCategory == constants.TPL_TREE {
		if v.TreeCode == "" {
			return errors.New("树编码不能为空")
		}
		if v.TreeName == "" {
			return errors.New("树名称不能为空")
		}
		if v.TreeParentCode == "" {
			return errors.New("树父编码不能为空")
		}
	} else if v.TplCategory == constants.TPL_SUB {
		if v.SubTableName == "" {
			return errors.New("子表名称不能为空")
		}
		if v.SubTableFkName == "" {
			return errors.New("子表外键名称不能为空")
		}
	}
	return nil
}
