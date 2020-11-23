package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client *mongo.Client
	err    error
)

type Student struct {
	Name string
	Age  int
}

func InitDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func InsertOne(s Student) {
	collection := client.Database("q1mi").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func UpdateOne() {
	collection := client.Database("q1mi").Collection("student")
	filter := bson.D{{
		"name", "zhouyang",
	}}
	update := bson.D{
		{
			"$inc", bson.D{{
				"age", 1,
			}},
		},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func DeleteOne() {
	collection := client.Database("q1mi").Collection("student")
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "zhouyang"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	// 删除所有
	deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)
}

func SelectOne() {
	var result Student
	filter := bson.D{{
		"name", "zhouyang",
	}}
	collection := client.Database("q1mi").Collection("student")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func main() {
	InitDB()
	student := Student{
		Name: "zhouyang",
		Age:  23,
	}
	InsertOne(student)
	SelectOne()
	//UpdateOne()
	//DeleteOne()
}
