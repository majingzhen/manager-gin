package constants

const (
	// TPL_CRUD 单表（增删改查）
	TPL_CRUD = "crud"
	// TPL_TREE 树表（增删改查）
	TPL_TREE = "tree"
	// TPL_SUB 主子表（增删改查）
	TPL_SUB = "sub"
	// TREE_CODE 树编码字段
	TREE_CODE = "treeCode"
	// TREE_PARENT_CODE 树父编码字段
	TREE_PARENT_CODE = "treeParentCode"
	// TREE_NAME 树名称字段
	TREE_NAME = "treeName"
	// PARENT_MENU_ID 上级菜单ID字段
	PARENT_MENU_ID = "parentMenuId"
	// PARENT_MENU_NAME 上级菜单名称字段
	PARENT_MENU_NAME = "parentMenuName"
	// HTML_INPUT 文本框
	HTML_INPUT = "input"

	// HTML_TEXTAREA 文本域
	HTML_TEXTAREA = "textarea"

	// HTML_SELECT 下拉框
	HTML_SELECT = "select"

	// HTML_RADIO 单选框
	HTML_RADIO = "radio"

	// HTML_CHECKBOX 复选框
	HTML_CHECKBOX = "checkbox"

	// HTML_DATETIME 日期控件
	HTML_DATETIME = "datetime"

	// HTML_IMAGE_UPLOAD 图片上传控件
	HTML_IMAGE_UPLOAD = "imageUpload"

	// HTML_FILE_UPLOAD 文件上传控件
	HTML_FILE_UPLOAD = "fileUpload"

	// HTML_EDITOR 富文本控件
	HTML_EDITOR = "editor"

	// TYPE_STRING 字符串类型
	TYPE_STRING = "string"

	// TYPE_INTERFACE 接口类型
	TYPE_INTERFACE = "interface{}"

	// 字符切片类型
	TYPE_BYTE_SLICE = "[]byte"

	// DEFAULT_STR 字符串默认值
	DEFAULT_STR = ""

	// DEFAULT_NUM 数字默认值
	DEFAULT_NUM = "0"

	// DEFAULT_BOOL 布尔值默认值
	DEFAULT_BOOL = "false"

	// DEFAULT_INTERFACE 接口默认值
	DEFAULT_INTERFACE = "nil"

	// TYPE_INTEGER 整型
	TYPE_INTEGER = "int"

	// TYPE_FLOAT 浮点型
	TYPE_FLOAT = "float64"

	// TYPE_DATE 时间类型
	TYPE_DATE = "*time.Time"

	// QUERY_LIKE 模糊查询
	QUERY_LIKE = "LIKE"

	// QUERY_EQ 相等查询
	QUERY_EQ = "EQ"

	// QUERY_BETWEEN 日期范围查询
	QUERY_BETWEEN = "BETWEEN"

	// REQUIRE 需要
	REQUIRE = "1"
)

var (
	// COLUMN_TYPE_STR  数据库字符串类型
	COLUMN_TYPE_STR = []string{"char", "varchar", "nvarchar", "varchar2"}
	// COLUMN_TYPE_TEXT 数据库文本类型
	COLUMN_TYPE_TEXT = []string{"tinytext", "text", "mediumtext", "longtext", "longblob"}
	// COLUMN_TYPE_TIME 数据库时间类型
	COLUMN_TYPE_TIME = []string{"datetime", "time", "date", "timestamp"}
	// COLUMN_TYPE_NUMBER 数据库数字类型
	COLUMN_TYPE_NUMBER = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer",
		"bit", "bigint"}
	// COLUMN_TYPE_FLOAT 数据库浮点类型
	COLUMN_TYPE_FLOAT = []string{"float", "double", "decimal"}
	// COLUMN_NAME_NOT_EDIT 页面不需要编辑字段
	COLUMN_NAME_NOT_EDIT = []string{"id", "create_by", "create_time", "is_del"}
	// COLUMN_NAME_NOT_LIST 页面不需要显示的列表字段
	COLUMN_NAME_NOT_LIST = []string{"id", "create_by", "create_time", "is_del", "update_by",
		"update_time"}

	// COLUMN_NAME_NOT_QUERY 页面不需要查询字段
	COLUMN_NAME_NOT_QUERY = []string{"id", "create_by", "create_time", "is_del", "update_by",
		"update_time", "remark"}

	// BASE_ENTITY Entity基类字段
	BASE_ENTITY = []string{"createBy", "createTime", "updateBy", "updateTime", "remark", "is_del"}

	// TREE_ENTITY Tree基类字段
	TREE_ENTITY = []string{"parentName", "parentId", "orderNum", "ancestors", "children"}

	DICT_HTML_TYPE = []string{HTML_SELECT, HTML_RADIO, HTML_CHECKBOX}
)
