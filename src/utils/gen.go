package utils

import (
	"manager-gin/src/app/admin/gen/model"
	"manager-gin/src/common/constants"
	"manager-gin/src/global"
	"strconv"
	"strings"
)

func InitTable(genTable *model.Table, operName string) {
	genTable.CreateBy = operName
	genTable.PackageName = global.GVA_VP.GetString("gen_code.package_name")
	genTable.ModuleName = genModuleName(global.GVA_VP.GetString("gen_code.package_name"))
	genTable.BusinessName = getBusinessName(genTable.Name)
	genTable.FunctionName = genTable.TableComment
	genTable.FunctionAuthor = global.GVA_VP.GetString("gen_code.author")
}

func InitColumnField(column *model.TableColumn, table *model.Table) {
	dataType := getColumnType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.Id
	column.CreateBy = table.CreateBy
	column.CreateTime = GetCurTime()
	// 设置结构体字段名
	column.GoField = ToCamelCase(columnName)
	// 设置默认类型
	column.GoType = constants.TYPE_STRING
	column.QueryType = constants.QUERY_EQ
	if Contains(constants.COLUMN_TYPE_STR, dataType) || Contains(constants.COLUMN_TYPE_TEXT, dataType) {
		columnLength := getColumnLength(column.ColumnType)
		if columnLength >= 500 || Contains(constants.COLUMN_TYPE_TEXT, dataType) {
			column.HtmlType = constants.HTML_TEXTAREA
		} else {
			column.HtmlType = constants.HTML_INPUT
		}
	} else if Contains(constants.COLUMN_TYPE_TIME, dataType) {
		column.HtmlType = constants.HTML_DATETIME
		column.GoType = constants.TYPE_DATE
	} else if Contains(constants.COLUMN_TYPE_NUMBER, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_INTEGER
	} else if Contains(constants.COLUMN_TYPE_FLOAT, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_FLOAT
	}
	// 插入字段
	column.IsInsert = constants.REQUIRE

	// 编辑字段
	if !Contains(constants.COLUMN_NAME_NOT_EDIT, columnName) && column.IsPk != "1" {
		column.IsEdit = constants.REQUIRE
	}
	// 列表字段
	if !Contains(constants.COLUMN_NAME_NOT_LIST, columnName) && column.IsPk != "1" {
		column.IsList = constants.REQUIRE
	}
	// 查询字段
	if !Contains(constants.COLUMN_NAME_NOT_QUERY, columnName) && column.IsPk != "1" {
		column.IsQuery = constants.REQUIRE
	}
	// 状态字段设置单选框
	if EndsWithIgnoreCase(columnName, "status") || EndsWithIgnoreCase(columnName, "flag") || BeginsWithIgnoreCase(columnName, "is") {
		column.HtmlType = constants.HTML_RADIO
	} else if EndsWithIgnoreCase(columnName, "type") || EndsWithIgnoreCase(columnName, "sex") {
		// 类型&性别字段设置下拉框
		column.HtmlType = constants.HTML_SELECT
	} else if EndsWithIgnoreCase(columnName, "image") {
		// 图片字段设置图片上传控件
		column.HtmlType = constants.HTML_IMAGE_UPLOAD
	} else if EndsWithIgnoreCase(columnName, "file") {
		// 文件字段设置文件上传控件
		column.HtmlType = constants.HTML_FILE_UPLOAD
	} else if EndsWithIgnoreCase(columnName, "content") {
		// 内容字段设置富文本控件
		column.HtmlType = constants.HTML_EDITOR
	}

	// 查询方式
	if column.IsQuery == constants.REQUIRE {
		if column.HtmlType == constants.HTML_DATETIME {
			column.QueryType = constants.QUERY_BETWEEN
		} else if column.HtmlType == constants.HTML_SELECT {
			column.QueryType = constants.QUERY_EQ
		} else {
			column.QueryType = constants.QUERY_LIKE
		}
	}

}

func getColumnLength(columnType string) int {
	if strings.Contains(columnType, "(") {
		length := columnType[strings.Index(columnType, "(")+1 : strings.Index(columnType, ")")]
		res, _ := strconv.Atoi(length)
		return res
	} else {
		return 0
	}
}

func getColumnType(columnType string) string {
	if strings.Contains(columnType, "(") {
		return columnType[:strings.Index(columnType, "(")-1]
	} else {
		return columnType
	}
}

func getBusinessName(name string) string {
	lastIndex := strings.LastIndex(name, "_")
	return name[lastIndex+1:]
}

func genModuleName(packageName string) string {
	lastIndex := strings.LastIndex(packageName, "/")
	return packageName[lastIndex+1:]
}
