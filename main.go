package main

import (
	"fmt"
	"github.com/mombe090/utils/mongo_utils"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const (
	databaseName   = "test"
	collectionName = "test-mongo_utils-go-driver"
)

type Test struct {
	Name string `json:"name"`
}

func testInsertOne() {
	res, err := mongo_utils.InsertOne(databaseName, collectionName, Test{Name: "test"})
	if err != nil {
		panic(err)
	}
	fmt.Sprintf(res)
}

func testInsertMany() {
	var testData []interface{}

	testData = append(testData, Test{Name: "Test 1"})
	testData = append(testData, Test{Name: "Test 2"})
	testData = append(testData, Test{Name: "Test 3"})

	res, err := mongo_utils.InsertMany(databaseName, collectionName, testData)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func testFindOne() {
	_, err := mongo_utils.FindOne(databaseName, collectionName, Test{Name: "Test 177"})
	if err != nil {
		panic(err)
	}

}

func testFindMany() {
	curr, ctx, err := mongo_utils.FindMany(databaseName, collectionName, Test{Name: "Test 1"})

	if err != nil {
		panic(err)
	}

	for curr.Next(*ctx) {
		var resTmp map[string]interface{}
		err := curr.Decode(&resTmp)

		if err != nil {
			log.Print(err)
		}
		fmt.Println(resTmp["_id"])
	}

	//log.Println("Finds", resultats)

}

func testUpdateOne() {
	err := mongo_utils.UpdateOne(databaseName, collectionName,
		bson.D{
			{
				"name", "test",
			},
		},
		bson.D{
			{"$set", bson.D{
				{
					"name", "tesssdfdsdfdsfdsfdst",
				},
			},
			},
		})

	if err != nil {
		panic(err)
	}
}

func testUpdateMany() {
	err := mongo_utils.UpdateOne(databaseName, collectionName,
		bson.D{
			{
				"name", "tesssdfdsdfdsfdsfdst",
			},
		},
		bson.D{
			{"$set", bson.D{
				{
					"name", "MoneyDoing",
				},
			},
			},
		})

	if err != nil {
		panic(err)
	}
}

func testDeleteOne() {
	err := mongo_utils.DeleteOne(
		databaseName,
		collectionName,
		bson.D{
			{
				"name", "MoneyDoing",
			},
		})

	if err != nil {
		panic(err)
	}
}

func testDeleteMany() {
	err := mongo_utils.DeleteMany(
		databaseName,
		collectionName,
		bson.D{
			{
				"name", "test",
			},
		})

	if err != nil {
		panic(err)
	}
}

func main() {
	//testInsertOne()
	//testInsertMany()
	testFindMany()
	testFindMany()
	//testFindOne()
	//testUpdateMany()
	//testDeleteOne()
	//testDeleteMany()

}
