package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"mombe090/utils/mongo_utils"
	"testing"
)

const (
	db   = "test"
	collection = "test-mongo_utils-go-driver"
	field = "from testing go test"
	fieldUpdate = "from testing go test update"
)

type TestingS struct {
	Val string `json:"value"`
}

func TestInsertAndDeleteOne(t *testing.T)  {
	err := mongo_utils.Insert(db, collection, TestingS{Val: field})
	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.Delete(db, collectionName, TestingS{Val:field})

	if err != nil {
		t.Error(err)
	}
}

func TestInsertAndDeleteMany(t *testing.T)  {
	var datas []interface{}

	datas = append(datas,
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
		)

	err := mongo_utils.Inserts(db, collection, datas)
	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.DeleteMany(db, collectionName, TestingS{Val:field + " go"})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOne(t *testing.T)  {
	err := mongo_utils.Insert(db, collection, TestingS{Val: field})

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.Update(db, collection, TestingS{Val:field},
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

	err = mongo_utils.Delete(db, collectionName, TestingS{Val:fieldUpdate})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateMany(t *testing.T)  {
	var datas []interface{}

	datas = append(datas,
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
		TestingS{Val: field + " go"},
	)

	err := mongo_utils.Inserts(db, collection, datas)

	if err != nil {
		t.Error(err)
	}

	err = mongo_utils.UpdateMany(db, collection, TestingS{Val:field + " go"},
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

	err = mongo_utils.DeleteMany(db, collectionName, TestingS{Val:fieldUpdate})
	if err != nil {
		t.Error(err)
	}
}
