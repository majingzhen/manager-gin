package utils

import (
	"manager-gin/src/app/admin/gen/model"
	"manager-gin/src/common/constants"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
	"strings"
)

// InitTable 初始化表结构体
func InitTable(genTable *model.Table, operName string) *model.Table {
	genTable.CreateBy = operName
	genTable.StructName = convertStructName(genTable.Name)
	genTable.PackageName = global.Viper.GetString("gen.package_name")
	genTable.ModuleName = genModuleName(global.Viper.GetString("gen.package_name"))
	genTable.BusinessName = getBusinessName(genTable.Name)
	genTable.FunctionName = genTable.TableComment
	genTable.FunctionAuthor = global.Viper.GetString("gen.author")
	return genTable
}

// convertStructName 转换struct名称
func convertStructName(tableName string) string {
	autoRemovePre := global.Viper.GetBool("gen.auto_remove_pre")
	tablePrefix := global.Viper.GetString("gen.table_prefix")
	if autoRemovePre && tablePrefix != "" {
		searchList := strings.Split(tablePrefix, ",")
		tableName = replaceFirst(tableName, searchList)
	}
	return utils.ToTitle(tableName)
}

// replaceFirst 批量替换前缀
func replaceFirst(replaceMen string, searchList []string) string {
	text := replaceMen
	for _, s := range searchList {
		if strings.HasPrefix(text, s) {
			text = strings.Replace(replaceMen, s, "", 1)
			break
		}
	}
	return text
}

// InitColumnField 初始化列属性字段
func InitColumnField(column *model.TableColumn, table *model.Table) *model.TableColumn {
	dataType := getColumnType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.Id
	column.CreateBy = table.CreateBy
	column.CreateTime = utils.GetCurTime()
	// 设置结构体字段名
	column.GoField = utils.ToCamelCase(columnName)
	// 设置默认类型
	column.GoType = constants.TYPE_STRING
	column.QueryType = constants.QUERY_EQ
	if utils.Contains(constants.COLUMN_TYPE_STR, dataType) || utils.Contains(constants.COLUMN_TYPE_TEXT, dataType) {
		columnLength := getColumnLength(column.ColumnType)
		if columnLength >= 500 || utils.Contains(constants.COLUMN_TYPE_TEXT, dataType) {
			column.HtmlType = constants.HTML_TEXTAREA
		} else {
			column.HtmlType = constants.HTML_INPUT
		}
	} else if utils.Contains(constants.COLUMN_TYPE_TIME, dataType) {
		column.HtmlType = constants.HTML_DATETIME
		column.GoType = constants.TYPE_DATE
	} else if utils.Contains(constants.COLUMN_TYPE_NUMBER, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_INTEGER
	} else if utils.Contains(constants.COLUMN_TYPE_FLOAT, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_FLOAT
	}
	// 插入字段
	column.IsInsert = constants.REQUIRE

	// 编辑字段
	if !utils.Contains(constants.COLUMN_NAME_NOT_EDIT, columnName) && column.IsPk != "1" {
		column.IsEdit = constants.REQUIRE
	}
	// 列表字段
	if !utils.Contains(constants.COLUMN_NAME_NOT_LIST, columnName) && column.IsPk != "1" {
		column.IsList = constants.REQUIRE
	}
	// 查询字段
	if !utils.Contains(constants.COLUMN_NAME_NOT_QUERY, columnName) && column.IsPk != "1" {
		column.IsQuery = constants.REQUIRE
	}
	// 状态字段设置单选框
	if utils.EndsWithIgnoreCase(columnName, "status") || utils.EndsWithIgnoreCase(columnName, "flag") || utils.BeginsWithIgnoreCase(columnName, "is") {
		column.HtmlType = constants.HTML_RADIO
	} else if utils.EndsWithIgnoreCase(columnName, "type") || utils.EndsWithIgnoreCase(columnName, "sex") {
		// 类型&性别字段设置下拉框
		column.HtmlType = constants.HTML_SELECT
	} else if utils.EndsWithIgnoreCase(columnName, "image") {
		// 图片字段设置图片上传控件
		column.HtmlType = constants.HTML_IMAGE_UPLOAD
	} else if utils.EndsWithIgnoreCase(columnName, "file") {
		// 文件字段设置文件上传控件
		column.HtmlType = constants.HTML_FILE_UPLOAD
	} else if utils.EndsWithIgnoreCase(columnName, "content") {
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
	return column
}

// getColumnByTableName 根据表名获取列信息
func getColumnLength(columnType string) int {
	if strings.Contains(columnType, "(") {
		length := columnType[strings.Index(columnType, "(")+1 : strings.Index(columnType, ")")]
		res, _ := strconv.Atoi(length)
		return res
	} else {
		return 0
	}
}

// getColumnType 获取列类型
func getColumnType(columnType string) string {
	if strings.Contains(columnType, "(") {
		return columnType[:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

// getBusinessName 获取业务名称
func getBusinessName(name string) string {
	lastIndex := strings.LastIndex(name, "_")
	return name[lastIndex+1:]
}

// genModuleName 获取模块名称
func genModuleName(packageName string) string {
	lastIndex := strings.LastIndex(packageName, "/")
	return packageName[lastIndex+1:]
}

// GenTemplatePath 获取模板路径
func GenTemplatePath(tplCategory string) []string {
	return []string{
		"./resources/tmpl/" + tplCategory + "/go/model.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/dao.go.txt",
		//"./resources/tmpl/" + tplCategory + "/go/service.txt",
		//"./resources/tmpl/" + tplCategory + "/go/view.txt",
		//"./resources/tmpl/" + tplCategory + "/go/view_utils.txt",
		//"./resources/tmpl/" + tplCategory + "/go/api.txt",
		//"./resources/tmpl/" + tplCategory + "/go/router.txt",
	}

}
