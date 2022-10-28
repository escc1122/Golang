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
)

func Test_facetGenerate(t *testing.T) {
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

	collection := client.Database("test").Collection("student")

	opt := &options.AggregateOptions{}

	matchOpt := pipeline.GetMatchGenerate().In("name", "escc1122", "escc1124").Eq("type", "A").GenBsonD()

	groupOpt := pipeline.GetGroupGenerate().GroupBy("name3", "name").GroupBy("type3", "type").Sum("sum_age", "age").Sum("sum_money", "money").Count("count").GenBsonD()

	facetGenerate := pipeline.GetFacetGenerate()
	groupPipeline := mongo.Pipeline{groupOpt}

	limitStage := bson.D{{"$limit", 1}}
	skipOpt := bson.D{{"$skip", 0}}
	sortOpt := bson.D{{"$sort", bson.M{"_id": 1}}}
	pipeline2 := mongo.Pipeline{sortOpt, skipOpt, limitStage}

	facetGenerate.AppendPipeline("groupPipeline", groupPipeline)
	facetGenerate.AppendPipeline("pipeline2", pipeline2)

	pipeline := mongo.Pipeline{matchOpt, facetGenerate.GenBsonD()}

	cursor, err := collection.Aggregate(context.TODO(), pipeline, opt)
	if err != nil {
		log.Println(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	//var strutResults []Student
	//
	//for cursor.Next(context.TODO()) {
	//	var result bson.M
	//	if err := cursor.Decode(&result); err != nil {
	//		log.Fatal(err)
	//	}
	//	results = append(results, result)
	//
	//	doc, _ := bson.Marshal(result)
	//	var student Student
	//	bson.Unmarshal(doc, &student)
	//	strutResults = append(strutResults, student)
	//	fmt.Println(result)
	//}
	//if err := cursor.Err(); err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(strutResults)
	fmt.Println(results)

	//data := results[0]["data"]

	//fmt.Println(data)

}
