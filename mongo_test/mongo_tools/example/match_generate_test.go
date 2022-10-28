package example

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongo_test/mongo_tools/pipeline"
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

func Test_matchGenerate_GteLt(t *testing.T) {
	opt := &options.AggregateOptions{}
	timeTest, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+08:00")

	matchOpt := pipeline.GetMatchGenerate().GteLt("time_test", timeTest.AddDate(0, 0, 1), timeTest.AddDate(0, 0, 2)).GenBsonD()
	matchPipeline := mongo.Pipeline{matchOpt}

	cursor, err := collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results)

	if len(results) != 1 {
		t.Errorf("len err")
	}

	if results[0]["_id"] != "id1" {
		t.Errorf("id err")
	}

	matchOpt = pipeline.GetMatchGenerate().In("name", "escc1122", "escc1124").Eq("type", "A").GteLt("age", 10, 40).GenBsonD()
	matchPipeline = mongo.Pipeline{matchOpt}
	cursor, err = collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results)

	if len(results) != 2 {
		t.Errorf("len err")
	}

}
