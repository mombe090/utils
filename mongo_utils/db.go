package mongo_utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
	"time"
)

const (
	dbhost    = "127.0.0.1:27017"
	authdb    = ""
	authuser  = ""
	authpass  = ""
	timeout   = 30 * time.Second
	poollimit = 4096
)

var client *mongo.Client
var ctx context.Context

func init() {
	ctx, _ = context.WithTimeout(context.Background(), timeout)
	//mongodb://user:password@host:port/userDb
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", dbhost)))

	if err != nil {
		log.Fatal(err)
	}

	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	client = c
}

func connect(databaseName string, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func Find(db string, col string, search interface{}) (interface{}, error) {
	collection := connect(db, col)

	var resultat interface{}

	err := collection.FindOne(ctx, search).Decode(&resultat)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			log.Println("No document founded {}")
			return resultat, nil
		}
		return nil, err
	}

	log.Println("Find", resultat)
	return resultat, nil
}

func FindMany(db string, col string, search interface{}) ([]interface{}, error) {
	collection := connect(db, col)

	var resultats []interface{}

	curr, err := collection.Find(ctx, search)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			log.Println("No document founded {}")
			return resultats, nil
		}
		return nil, err
	}

	for curr.Next(ctx) {
		var resTmp bson.M
		err := curr.Decode(&resTmp)

		if err != nil {
			log.Print(err)
			return nil, err
		}
		resultats = append(resultats, resTmp)
	}

	log.Println("Finds", resultats)
	return resultats, nil
}

func Insert(db string, col string, data interface{}) error {
	collection := connect(db, col)

	res, errInsert := collection.InsertOne(ctx, data)

	if errInsert != nil {
		return errInsert
	}

	log.Println(res.InsertedID)
	return nil
}

func Inserts(db string, col string, data []interface{}) error {
	collection := connect(db, col)

	res, errInserts := collection.InsertMany(ctx, data)

	if errInserts != nil {
		return errInserts
	}

	log.Println("Successfully add with IDS ", res.InsertedIDs)
	return nil
}

func Update(db string, col string, search interface{}, data interface{}) error {
	collection := connect(db, col)

	res, errInsert := collection.UpdateOne(ctx, search, data)

	if errInsert != nil {
		return errInsert
	}

	log.Println("Matched", res.MatchedCount)
	log.Println("Modified", res.ModifiedCount)
	return nil
}

func UpdateMany(db string, col string, search interface{}, data interface{}) error {
	collection := connect(db, col)

	res, errInsert := collection.UpdateMany(ctx, search, data)

	if errInsert != nil {
		return errInsert
	}

	log.Println("Matched", res.MatchedCount)
	log.Println("Modified", res.ModifiedCount)
	return nil
}

func Delete(db string, col string, search interface{}) error {
	collection := connect(db, col)

	res, errInsert := collection.DeleteOne(ctx, search)

	if errInsert != nil {
		return errInsert
	}

	log.Println("Deleted ", res.DeletedCount)
	return nil
}

func DeleteMany(db string, col string, search interface{}) error {
	collection := connect(db, col)

	res, errInsert := collection.DeleteMany(ctx, search)

	if errInsert != nil {
		return errInsert
	}

	log.Println("Deleted ", res.DeletedCount)
	return nil
}
