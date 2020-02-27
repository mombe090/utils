package main

import (
	"github.com/mombe090/utils/mongo_utils"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

const (
	db          = "test"
	collection  = "test-mongo_utils-go-driver"
	field       = "from testing go test"
	fieldUpdate = "from testing go test update"
)

type TestingS struct {
	Val string `json:"value"`
}

func TestInsertAndDeleteOne(t *testing.T) {
	_, err := mongo_utils.InsertOne(db, collection, TestingS{Val: field})
	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.DeleteOne(db, collectionName, TestingS{Val: field})

	if err != nil {
		t.Error(err)
	}
}

func TestInsertAndDeleteMany(t *testing.T) {
	var datas []interface{}

	datas = append(datas,
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
	)

	_, err := mongo_utils.InsertMany(db, collection, datas)
	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.DeleteMany(db, collectionName, TestingS{Val: field + " go"})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOne(t *testing.T) {
	_,  err := mongo_utils.InsertOne(db, collection, TestingS{Val: field})

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.UpdateOne(db, collection, TestingS{Val: field},
		bson.D{
			{"$set", bson.D{
				{
					"name", fieldUpdate,
				},
			},
			},
		})

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.DeleteOne(db, collectionName, TestingS{Val: fieldUpdate})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateMany(t *testing.T) {
	var datas []interface{}

	datas = append(datas,
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
	)

	_, err := mongo_utils.InsertMany(db, collection, datas)

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.UpdateMany(db, collection, TestingS{Val: field + " go"},
		bson.D{
			{"$set", bson.D{
				{
					"name", fieldUpdate,
				},
			},
			},
		})

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.DeleteMany(db, collectionName, TestingS{Val: fieldUpdate})
	if err != nil {
		t.Error(err)
	}
}
