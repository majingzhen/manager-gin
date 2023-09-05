package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GormDao     *gorm.DB
	DbList      map[string]*gorm.DB
	RedisClient *redis.Client
	Viper       *viper.Viper
	Logger      *zap.Logger
	//GVA_Concurrency_Control = &singleflight.Group{}
	//BlackCache              local_cache.Cache
	lock        sync.RWMutex
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DbList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DbList[dbname]
	if !ok || db == nil {
		panic("db no before")
	}
	return db
}
