package mongo_tools

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

type UpdateFilterTestSuit struct {
	suite.Suite
	testCount  uint32
	collection *mongo.Collection
}

func (s *UpdateFilterTestSuit) SetupSuite() {
	fmt.Println("SetupSuite")

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

	s.collection = client.Database("test").Collection("student")
}

func (s *UpdateFilterTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *UpdateFilterTestSuit) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", s.testCount)

	time_test, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+08:00")

	s1 := Student{
		ID:       "id1",
		Name:     "escc1122",
		Type:     "A",
		Age:      33,
		Money:    124,
		TimeTest: time_test.AddDate(0, 0, 1),
	}

	s2 := Student{
		ID:       "id2",
		Name:     "escc1122",
		Type:     "A",
		Age:      55,
		Money:    124,
		TimeTest: time_test.AddDate(0, 0, 2),
	}
	s3 := Student{
		ID:       "id3",
		Name:     "escc1124",
		Type:     "A",
		Age:      23,
		Money:    234,
		TimeTest: time_test.AddDate(0, 0, 3),
	}

	s4 := Student{
		ID:       "id4",
		Name:     "escc1124",
		Type:     "B",
		Age:      44,
		Money:    34535,
		TimeTest: time_test.AddDate(0, 0, 4),
	}

	s5 := Student{
		ID:       "id5",
		Name:     "escc1122",
		Type:     "B",
		Age:      44,
		Money:    234,
		TimeTest: time_test.AddDate(0, 0, 5),
	}

	s6 := Student{
		ID:       "id6",
		Name:     "escc1122",
		Type:     "C",
		Age:      4,
		Money:    345,
		TimeTest: time_test.AddDate(0, 0, 6),
	}

	sArr := []interface{}{
		s2, s1, s3, s4, s5, s6,
	}

	//insertResult, err := s.collection.InsertOne(context.TODO(), s1)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(insertResult)

	opt := &options.InsertManyOptions{}
	//opt.SetOrdered(false)
	opt.SetOrdered(false)

	_, err2 := s.collection.InsertMany(context.TODO(), sArr, opt)
	if err2 != nil {
		log.Println(err2)
	}

	//fmt.Println(insertResult2)
	//
	//bulkOption := options.BulkWriteOptions{}
	//bulkOption.SetOrdered(false)
	//var operations []mongo.WriteModel
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s1))
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s2))
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s3))
	//
	//bulkWrite, err3 := s.collection.BulkWrite(context.TODO(), operations, &bulkOption)
	//
	//if err3 != nil {
	//	log.Println(err3)
	//}
	//
	//if reflect.TypeOf(err3).Name() == "BulkWriteException" {
	//	fmt.Println("test")
	//}
	//
	//fmt.Println(bulkWrite)
}

func (s *UpdateFilterTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
	s.collection.Drop(context.TODO())
}

func (s *UpdateFilterTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *UpdateFilterTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *UpdateFilterTestSuit) TestUpdate() {
	filter := GetFilterGenerate().Eq("_id", "id6").GenBsonD()

	fields := map[string]interface{}{}
	fields["name"] = "test update"
	fields["age"] = 55

	s.collection.UpdateMany(context.TODO(), filter, GenSimpleBsonD("$set", GenBsonD(fields)))

	matchGenerate := GetMatchGenerate().Eq("_id", "id6").GenBsonD()

	opt := &options.AggregateOptions{}

	var results []*Student
	matchPipeline := mongo.Pipeline{matchGenerate}
	cursor, err := s.collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	assert.Equal(s.T(), 1, len(results))
	assert.Equal(s.T(), "id6", results[0].ID)
	assert.Equal(s.T(), 55, results[0].Age)
	assert.Equal(s.T(), "C", results[0].Type)
	assert.Equal(s.T(), 345, results[0].Money)

}

type FilterTestSuit struct {
	suite.Suite
	testCount  uint32
	collection *mongo.Collection
}

func (s *FilterTestSuit) SetupSuite() {
	fmt.Println("SetupSuite")

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

	s.collection = client.Database("test").Collection("student")
}

func (s *FilterTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *FilterTestSuit) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", s.testCount)

	time_test, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+08:00")

	s1 := Student{
		ID:       "id1",
		Name:     "escc1122",
		Type:     "A",
		Age:      33,
		Money:    124,
		TimeTest: time_test.AddDate(0, 0, 1),
	}

	s2 := Student{
		ID:       "id2",
		Name:     "escc1122",
		Type:     "A",
		Age:      55,
		Money:    124,
		TimeTest: time_test.AddDate(0, 0, 2),
	}
	s3 := Student{
		ID:       "id3",
		Name:     "escc1124",
		Type:     "A",
		Age:      23,
		Money:    234,
		TimeTest: time_test.AddDate(0, 0, 3),
	}

	s4 := Student{
		ID:       "id4",
		Name:     "escc1124",
		Type:     "B",
		Age:      44,
		Money:    34535,
		TimeTest: time_test.AddDate(0, 0, 4),
	}

	s5 := Student{
		ID:       "id5",
		Name:     "escc1122",
		Type:     "B",
		Age:      44,
		Money:    234,
		TimeTest: time_test.AddDate(0, 0, 5),
	}

	s6 := Student{
		ID:       "id6",
		Name:     "escc1122",
		Type:     "C",
		Age:      4,
		Money:    345,
		TimeTest: time_test.AddDate(0, 0, 6),
	}

	sArr := []interface{}{
		s2, s1, s3, s4, s5, s6,
	}

	//insertResult, err := s.collection.InsertOne(context.TODO(), s1)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(insertResult)

	opt := &options.InsertManyOptions{}
	//opt.SetOrdered(false)
	opt.SetOrdered(false)

	_, err2 := s.collection.InsertMany(context.TODO(), sArr, opt)
	if err2 != nil {
		log.Println(err2)
	}

	//fmt.Println(insertResult2)
	//
	//bulkOption := options.BulkWriteOptions{}
	//bulkOption.SetOrdered(false)
	//var operations []mongo.WriteModel
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s1))
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s2))
	//operations = append(operations, mongo.NewInsertOneModel().SetDocument(s3))
	//
	//bulkWrite, err3 := s.collection.BulkWrite(context.TODO(), operations, &bulkOption)
	//
	//if err3 != nil {
	//	log.Println(err3)
	//}
	//
	//if reflect.TypeOf(err3).Name() == "BulkWriteException" {
	//	fmt.Println("test")
	//}
	//
	//fmt.Println(bulkWrite)
}

func (s *FilterTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
	s.collection.Drop(context.TODO())
}

func (s *FilterTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *FilterTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *FilterTestSuit) TestFilterIn() {
	filterGenerate := GetFilterGenerate().In("_id", "id1", "id2").GenBsonD()

	cursor, err := s.collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})

	var results []*Student
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	assert.Equal(s.T(), 2, len(results))
	assert.Equal(s.T(), "id1", results[0].ID)
	assert.Equal(s.T(), "id2", results[1].ID)
}

func (s *FilterTestSuit) TestFilterEq() {
	filterGenerate := GetFilterGenerate().Eq("_id", "id1").GenBsonD()
	cursor, err := s.collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})
	var results []*Student
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	assert.Equal(s.T(), 1, len(results))
	assert.Equal(s.T(), "id1", results[0].ID)
}

func (s *FilterTestSuit) TestFilterGteLt() {
	timeTest, _ := time.Parse(time.RFC3339, "2020-04-01T00:00:00+08:00")

	filterGenerate := GetFilterGenerate().GteLt("time_test", timeTest.AddDate(0, 0, 1), timeTest.AddDate(0, 0, 2)).GenBsonD()

	cursor, err := s.collection.Find(context.TODO(), filterGenerate, &options.FindOptions{})
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	assert.Equal(s.T(), 1, len(results))
	assert.Equal(s.T(), "id1", results[0]["_id"])
}

func TestFilter(t *testing.T) {
	suite.Run(t, new(UpdateFilterTestSuit))
	suite.Run(t, new(FilterTestSuit))
}
