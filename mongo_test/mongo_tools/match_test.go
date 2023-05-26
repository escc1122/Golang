package mongo_tools

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

type MatchTestSuit struct {
	suite.Suite
	testCount  uint32
	collection *mongo.Collection
}

func (s *MatchTestSuit) SetupSuite() {
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

func (s *MatchTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *MatchTestSuit) SetupTest() {
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

	opt := &options.InsertManyOptions{}
	opt.SetOrdered(false)

	_, err2 := s.collection.InsertMany(context.TODO(), sArr, opt)
	if err2 != nil {
		log.Println(err2)
	}
}

func (s *MatchTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
	s.collection.Drop(context.TODO())
}

func (s *MatchTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MatchTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MatchTestSuit) TestMatchGteLt() {
	opt := &options.AggregateOptions{}
	var results []*Student
	matchOpt := GetMatchGenerate().GteLt("age", 10, 40).GenBsonD()
	matchPipeline := mongo.Pipeline{matchOpt}
	cursor, err := s.collection.Aggregate(context.TODO(), matchPipeline, opt)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	assert.Equal(s.T(), 2, len(results))
	assert.Equal(s.T(), "id1", results[0].ID)
	assert.Equal(s.T(), "id3", results[1].ID)
}

func TestMatch(t *testing.T) {
	suite.Run(t, new(MatchTestSuit))
}
