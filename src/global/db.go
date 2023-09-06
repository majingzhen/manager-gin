package global

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitDataSource 初始化数据库
func InitDataSource() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		Viper.GetString("datasource.username"),
		Viper.GetString("datasource.password"),
		Viper.GetString("datasource.host"),
		Viper.GetString("datasource.port"),
		Viper.GetString("datasource.db_name"))
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Viper.GetString("database.table_prefix"), // 控制表前缀
			SingularTable: true,
		},
		Logger: logger.Default, // 控制是否sql输出，默认是不输出
	}
	if Viper.GetBool("database.log_mode") {
		gcf.Logger = logger.Default.LogMode(logger.Info) // logger.Info 就会输出sql
	}

	if tmp, err := gorm.Open(mysql.Open(dsn), gcf); err != nil {
		Logger.Error("MySQL启动异常", zap.Error(err))
		panic(err)
	} else {
		//Logger.Info("Connect to database success")
		//// 全局禁用表名复数
		//tmp = tmp.Set("gorm:table_options", "ENGINE=InnoDB")
		//// 全局设置表前缀
		//sqlDB, _ := tmp.DB()
		//sqlDB.SetMaxIdleConns(10)
		//sqlDB.SetMaxOpenConns(100)
		//sqlDB.SetConnMaxLifetime(10)
		//DbList = make(map[string]*gorm.DB)
		//DbList[Viper.GetString("datasource.db_name")] = GormDao
		GormDao = tmp
	}
}
