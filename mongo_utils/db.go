package mongo_utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
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
	gotenv.Load()
	ctx, _ = context.WithTimeout(context.Background(), timeout)
	//mongodb://user:password@host:port/userDb
	var c *mongo.Client
	var err error

	if os.Getenv("MONGOHOST") != "" && os.Getenv("MONGOPORT") != ""{
		if os.Getenv("MONGOUSER") != "" && os.Getenv("MONGOPASSWORD") != "" && os.Getenv("MONGOAUTHDB") != "" {
			c, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",os.Getenv("MONGOUSER"), os.Getenv("MONGOPASSWORD"), os.Getenv("MONGOHOST"), os.Getenv("MONGOPORT"), os.Getenv("MONGOAUTHDB"))))
		} else {
			c, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGOHOST"), os.Getenv("MONGOPORT"))))
		}
	} else {
		c, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", dbhost)))
	}


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

func FindOne(db string, col string, search interface{}) (interface{}, error) {
	collection := connect(db, col)

	var resultat map[string]interface{}

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

func FindMany(db string, col string, search interface{}) (*mongo.Cursor, *context.Context, error) {
	collection := connect(db, col)

	curr, err := collection.Find(ctx, search)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			log.Println("No document founded {}")
			return nil, &ctx, nil
		}
		return nil, &ctx, err
	}

	return curr, &ctx, nil
}

func InsertOne(db string, col string, data interface{}) (string, error) {
	collection := connect(db, col)

	res, errInsert := collection.InsertOne(ctx, data)

	if errInsert != nil {
		return "", errInsert
	}


	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	} else {
		return "", errors.New("not objectId returned")
	}
}

func InsertMany(db string, col string, data []interface{}) ([]string, error) {
	collection := connect(db, col)

	res, errInserts := collection.InsertMany(ctx, data)

	if errInserts != nil {
		return nil, errInserts
	}

	var ids []string
	for _, e := range res.InsertedIDs {
		if oid, ok := e.(primitive.ObjectID); ok {
			ids = append(ids, oid.Hex())
		}
	}

	log.Println("Successfully add with IDS ", res.InsertedIDs)
	return ids, nil
}

func UpdateOne(db string, col string, search interface{}, data interface{}) error {
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

func DeleteOne(db string, col string, search interface{}) error {
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
