package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		GVA_VP.GetString("database.username"),
		GVA_VP.GetString("database.password"),
		GVA_VP.GetString("database.host"),
		GVA_VP.GetString("database.port"),
		GVA_VP.GetString("database.db_name"))
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   GVA_VP.GetString("database.table_prefix"), // 控制表前缀
			SingularTable: true,
		},
		Logger: logger.Default, // 控制是否sql输出，默认是不输出
	}
	if GVA_VP.GetBool("database.log_mode") {
		gcf.Logger = logger.Default.LogMode(logger.Info) // logger.Info 就会输出sql
	}
	GOrmDao, _ = gorm.Open(mysql.Open(dsn), gcf)
}
