package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongo_test/mongo_tools/pipeline"
)

type Student struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Type  string `bson:"type"`
	Money int    `bson:"money"`
}

func main() {
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

	testInsert(collection)

	testAgg(collection)

}

func testInsert(collection *mongo.Collection) {
	s1 := Student{
		ID:    "id1",
		Name:  "escc1122",
		Type:  "A",
		Age:   33,
		Money: 124,
	}

	s2 := Student{
		ID:    "id2",
		Name:  "escc1122",
		Type:  "A",
		Age:   55,
		Money: 124,
	}
	s3 := Student{
		ID:    "id3",
		Name:  "escc1124",
		Type:  "A",
		Age:   23,
		Money: 234,
	}

	s4 := Student{
		ID:    "id4",
		Name:  "escc1124",
		Type:  "B",
		Age:   44,
		Money: 34535,
	}

	s5 := Student{
		ID:    "id5",
		Name:  "escc1122",
		Type:  "B",
		Age:   44,
		Money: 234,
	}

	s6 := Student{
		ID:    "id6",
		Name:  "escc1122",
		Type:  "C",
		Age:   4,
		Money: 345,
	}

	sArr := []interface{}{
		s2, s1, s3, s4, s5, s6,
	}

	insertResult, err := collection.InsertOne(context.TODO(), s1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(insertResult)

	opt := &options.InsertManyOptions{}
	//opt.SetOrdered(false)
	opt.SetOrdered(false)

	insertResult2, err2 := collection.InsertMany(context.TODO(), sArr, opt)
	if err2 != nil {
		log.Println(err2)
	}

	fmt.Println(insertResult2)

	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(false)
	var operations []mongo.WriteModel
	operations = append(operations, mongo.NewInsertOneModel().SetDocument(s1))
	operations = append(operations, mongo.NewInsertOneModel().SetDocument(s2))
	operations = append(operations, mongo.NewInsertOneModel().SetDocument(s3))

	bulkWrite, err3 := collection.BulkWrite(context.TODO(), operations, &bulkOption)

	if err3 != nil {
		log.Println(err2)
	}
	fmt.Println(bulkWrite)
}

func testAgg(collection *mongo.Collection) {
	opt := &options.AggregateOptions{
		AllowDiskUse:             nil,
		BatchSize:                nil,
		BypassDocumentValidation: nil,
		Collation:                nil,
		MaxTime:                  nil,
		MaxAwaitTime:             nil,
		Comment:                  nil,
		Hint:                     nil,
		Let:                      nil,
	}
	//pipeline := []bson.M{
	//	{"$group": bson.M{
	//		//"_id":  "$name",
	//		//"_id":  bson.D{{"name", "$name"}, {"id", "$_id"}},
	//		//"_id":   "$_id",
	//		"_id":  "$name",
	//		"_sun": bson.D{{"$sum", "$age"}},
	//	}},
	//	{"$match": bson.D{{"$name", "escc1122"}}},
	//}

	groupOpt2 := bson.D{
		{"$group", bson.M{
			//"_id":  "$name",
			"_id": bson.D{{"name", "$name"}, {"type", "$type"}},
			//"_id":   "$_id",
			//"_id":  "$name",
			"_sun": bson.D{{"$sum", "$age"}},
		}},
	}

	matchOpt2 := bson.D{
		{"$match", bson.M{
			"type": bson.D{{"$eq", "B"}},
			"name": bson.D{{"$in", bson.A{"escc1122", "escc1124"}}},
		}},
	}
	fmt.Println(groupOpt2)
	fmt.Println(matchOpt2)

	matchOpt := pipeline.GetMatchOpt().In("name", "escc1122", "escc1124").Gen()

	//matchOpt := getMatchOpt().eq("name", "escc1122").gen()

	//matchOpt := mongo_tools.GetMatchOpt().In("name", "escc1122", "escc1124").Eq("type", "B").Gen()

	//matchOpt := mongo_tools.GetMatchOpt().Gt("age", 40).Lt("age", 50).Gen()
	//matchOpt := mongo_tools.GetMatchOpt().Gte("age", 40).Lte("age", 50).Gen()
	//matchOpt := mongo_tools.GetMatchOpt().GteLt("age", 40, 50).Gen()

	groupOpt := pipeline.GetGroupOpt().SetGroupPara("name", "name").SetGroupPara("type", "type").SetSum("sum_age", "age").SetSum("sum_money", "money").Gen()

	//limitStage := bson.D{{"$limit", 2}}

	//sortStage := bson.D{{"$sort", bson.M{"_id": 1}}}

	fmt.Println(groupOpt)
	fmt.Println(matchOpt)
	//fmt.Println(sortStage)

	sortOpt := pipeline.GetSortOpt().SetSort("name", -1).SetSort("_id", -1).Gen()

	pipeline := mongo.Pipeline{matchOpt, sortOpt, groupOpt}

	result, err := collection.Aggregate(context.TODO(), pipeline, opt)
	if err != nil {
		log.Println(err)
	}

	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results)

}
