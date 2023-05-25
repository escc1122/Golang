package mongo_tools

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"reflect"
	"testing"
	"time"
)

type Student struct {
	ID       string    `bson:"_id"`
	Name     string    `bson:"name"`
	Age      int       `bson:"age"`
	Type     string    `bson:"type"`
	Money    int       `bson:"money"`
	TimeTest time.Time `bson:"time_test"`
}

type BucketTestSuit struct {
	suite.Suite
	testCount  uint32
	collection *mongo.Collection
}

func (s *BucketTestSuit) SetupSuite() {
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

func (s *BucketTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *BucketTestSuit) SetupTest() {
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

	insertResult, err := s.collection.InsertOne(context.TODO(), s1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(insertResult)

	opt := &options.InsertManyOptions{}
	//opt.SetOrdered(false)
	opt.SetOrdered(false)

	insertResult2, err2 := s.collection.InsertMany(context.TODO(), sArr, opt)
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

	bulkWrite, err3 := s.collection.BulkWrite(context.TODO(), operations, &bulkOption)

	if err3 != nil {
		log.Println(err3)
	}

	if reflect.TypeOf(err3).Name() == "BulkWriteException" {
		fmt.Println("test")
	}

	fmt.Println(bulkWrite)
}

func (s *BucketTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
	s.collection.Drop(context.TODO())
}

func (s *BucketTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *BucketTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *BucketTestSuit) TestBucket() {
	fmt.Println("TestExample")

	int64Max := math.MaxInt64
	int32Max := math.MaxInt
	int64Min := math.MinInt64
	fmt.Println(int64Min)
	if int64Max == int32Max {

	}

	a := []interface{}{0, 2, 4, 10, 15}

	bucketGenerate := GetBucketGenerate("age", a, int64Max)

	bucketGenerate.Count("count")

	bucketGenerate.Sum("sum_test", "age")

	bucketGenerate.Max("max_age", "age")

	bbb := bucketGenerate.GenBsonD()

	matchPipeline := mongo.Pipeline{bbb}

	opt := &options.AggregateOptions{}
	cursor, err := s.collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}

	type resp struct {
		ID      int64 `bson:"_id"`
		Count   int64 `bson:"count"`
		SumTest int64 `bson:"sum_test"`
		MaxAge  int64 `bson:"max_age"`
	}

	var results []*resp
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results)

	assert.Equal(s.T(), 2, len(results))
	assert.Equal(s.T(), int64(4), results[0].ID)
	assert.Equal(s.T(), int64(1), results[0].Count)
	assert.Equal(s.T(), int64(4), results[0].SumTest)
	assert.Equal(s.T(), int64(4), results[0].MaxAge)

	assert.Equal(s.T(), int64(9223372036854775807), results[1].ID)
	assert.Equal(s.T(), int64(5), results[1].Count)
	assert.Equal(s.T(), int64(199), results[1].SumTest)
	assert.Equal(s.T(), int64(55), results[1].MaxAge)

}

func TestExample(t *testing.T) {
	suite.Run(t, new(BucketTestSuit))
}
