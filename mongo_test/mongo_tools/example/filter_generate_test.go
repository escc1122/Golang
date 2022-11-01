package example

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongo_test/mongo_tools"
	"testing"
	"time"
)

var collection *mongo.Collection

func init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("test").Collection("student")
}

func Test_filterGenerate_In(t *testing.T) {
	filterGenerate := mongo_tools.GetFilterGenerate().In("_id", "id1", "id2").GenBsonD()

	cursor, err := collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) != 2 {
		t.Errorf("len err")
	}
	if results[0]["_id"] != "id1" {
		t.Errorf("id err")
	}
	if results[1]["_id"] != "id2" {
		t.Errorf("id err")
	}
}

func Test_filterGenerate_Eq(t *testing.T) {
	filterGenerate := mongo_tools.GetFilterGenerate().Eq("_id", "id1").GenBsonD()
	cursor, err := collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) != 1 {
		t.Errorf("len err")
	}
	if results[0]["_id"] != "id1" {
		t.Errorf("id err")
	}
}

func Test_filterGenerate_GteLt(t *testing.T) {

	timeTest, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+08:00")

	filterGenerate := mongo_tools.GetFilterGenerate().GteLt("time_test", timeTest.AddDate(0, 0, 1), timeTest.AddDate(0, 0, 2)).GenBsonD()

	cursor, err := collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) != 1 {
		t.Errorf("len err")
	}

	if results[0]["_id"] != "id1" {
		t.Errorf("id err")
	}

	filterGenerate = mongo_tools.GetFilterGenerate().In("name", "escc1122", "escc1124").Eq("type", "A").GteLt("age", 10, 40).GenBsonD()

	cursor, err = collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})

	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) != 2 {
		t.Errorf("len err")
	}

}
