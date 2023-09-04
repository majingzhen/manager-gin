// Package table_column 自动生成模板 TableColumnService
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package table_column

import (
	"manager-gin/src/app/admin/gen/dao"
	"manager-gin/src/app/admin/gen/service/table_column/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

type TableColumnService struct {
	tableColumnDao dao.TableColumnDao
	viewUtils      view.TableColumnViewUtils
}

// DeleteByIds 批量删除TableColumn记录
// Author
func (s *TableColumnService) DeleteByIds(ids []string) (err error) {
	err = s.tableColumnDao.DeleteByIds(global.GOrmDao, ids)
	return err
}

// Update 更新TableColumn记录
// Author
func (s *TableColumnService) Update(id string, tableColumnView *view.TableColumnView) error {
	tableColumnView.Id = id
	if err, tableColumn := s.viewUtils.View2Data(tableColumnView); err != nil {
		return err
	} else {
		return s.tableColumnDao.Update(*tableColumn)
	}
}

// Get 根据id获取TableColumn记录
// Author
func (s *TableColumnService) Get(id string) (err error, tableColumnView *view.TableColumnView) {
	err1, tableColumn := s.tableColumnDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, tableColumnView = s.viewUtils.Data2View(tableColumn)
	return
}

// Page 分页获取TableColumn记录
// Author
func (s *TableColumnService) Page(pageInfo *view.TableColumnPageView) (err error, res *common.PageInfo) {
	if err, res = s.tableColumnDao.Page(pageInfo); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// List 获取TableColumn列表
// Author
func (s *TableColumnService) List(v *view.TableColumnQueryView) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.List(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

// GetColumnListByTableId 根据tableId获取TableColumn列表
func (s *TableColumnService) GetColumnListByTableId(tableId string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.GetColumnListByTableId(tableId); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}
