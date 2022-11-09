package example

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"mongo_test/mongo_tools"
	"testing"
)

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

func Test_bucket(t *testing.T) {
	int64Max := math.MaxInt64
	int32Max := math.MaxInt
	int64Min := math.MinInt64
	fmt.Println(int64Min)
	if int64Max == int32Max {

	}

	a := []interface{}{0, 2, 4, 10, 15}

	bucketGenerate := mongo_tools.GetBucketGenerate("age", a, int64Max)

	bucketGenerate.Count("count")

	bucketGenerate.Sum("sum_test", "age")

	bucketGenerate.Max("max_age", "age")

	//bucketGenerate.GroupBy("age")
	//bucketGenerate.Count("count")
	//bucketGenerate.Set()

	//a := []interface{}{20, 50}

	//bucketGenerate.Boundaries(a)

	bbb := bucketGenerate.GenBsonD()

	matchPipeline := mongo.Pipeline{bbb}

	opt := &options.AggregateOptions{}
	cursor, err := collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results)

	//if len(results) != 1 {
	//	t.Errorf("len err")
	//}
	//
	//if results[0]["_id"] != "id1" {
	//	t.Errorf("id err")
	//}
}
