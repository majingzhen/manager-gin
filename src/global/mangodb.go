package global

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// InitMongoClient 初始化全局mongo客户端
func InitMongoClient() {
	uri := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", Viper.GetString("datasource.mongodb.host"), Viper.GetString("datasource.mongodb.port")))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		Logger.Error("mongo connect error", zap.Error(err))
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Error("mongo ping error", zap.Error(err))
		panic(err)
	}
	Logger.Info("Connected to MongoDB!")
	MongoClient = client
	MongoDB = MongoClient.Database(Viper.GetString("datasource.mongodb.db_name"))
}

// CloseMongoClient 关闭全局mongo客户端
func CloseMongoClient() {
	if MongoClient != nil {
		if err := MongoClient.Disconnect(context.Background()); err != nil {
			Logger.Error("关闭mongodb错误")
		}
	}
}
