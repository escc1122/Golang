package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Student struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
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

	s1 := Student{
		ID:   "dsfdsfd",
		Name: "111edsfsdfds",
		Age:  5655,
	}

	s2 := Student{
		ID:   "346436464",
		Name: "3253253",
		Age:  12442,
	}
	s3 := Student{
		ID:   "346fff436464",
		Name: "3253aa253",
		Age:  23,
	}

	sArr := []interface{}{
		s2, s1, s3,
	}

	collection := client.Database("test").Collection("student")

	insertResult, err := collection.InsertOne(context.TODO(), s1)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(insertResult)

	opt := &options.InsertManyOptions{}
	//opt.SetOrdered(false)
	opt.SetOrdered(true)

	insertResult2, err2 := collection.InsertMany(context.TODO(), sArr, opt)
	if err2 != nil {
		log.Println(err2)
	}

	fmt.Println(insertResult2)
}
