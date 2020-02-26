package main

import (
	"github.com/mombe090/utils/mongo_utils"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	databaseName   = "test"
	collectionName = "test-mongo_utils-go-driver"
)

type Test struct {
	Name string `json:"name"`
}

func testInsertOne() {
	err := mongo_utils.Insert(databaseName, collectionName, Test{Name: "test"})
	if err != nil {
		panic(err)
	}
}

func testInsertMany() {
	var testData []interface{}

	testData = append(testData, Test{Name: "Test 1"})
	testData = append(testData, Test{Name: "Test 2"})
	testData = append(testData, Test{Name: "Test 3"})

	err := mongo_utils.Inserts(databaseName, collectionName, testData)
	if err != nil {
		panic(err)
	}
}

func testFindOne() {
	_, err := mongo_utils.Find(databaseName, collectionName, Test{Name: "Test 177"})
	if err != nil {
		panic(err)
	}

}

func testFindMany() {
	_, err := mongo_utils.FindMany(databaseName, collectionName, Test{Name: "Test 1"})
	if err != nil {
		panic(err)
	}

}

func testUpdateOne() {
	err := mongo_utils.Update(databaseName, collectionName,
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
	err := mongo_utils.Update(databaseName, collectionName,
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
	err := mongo_utils.Delete(
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
	testInsertOne()
	//testInsertMany()
	//testFindOne()
	//testUpdateMany()
	//testDeleteOne()
	testDeleteMany()

}
