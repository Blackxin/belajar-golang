package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}

func main() {
	// A.57.2. Insert Data
	// insert()

	// A.57.3. Membaca Data
	// find()

	// A.57.4. Update Data
	// update()

	// A.57.5. Menghapus Data
	// remove()

	// A.57.6. Aggregate Data
	aggregate()
}

// utils
func connect() (*mongo.Database, error) {
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb://%s:%s@localhost:27017", dbUser, dbPwd))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Wick", 2})
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("insert success!")
}

func find() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	if err != nil {
		panic(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name\t:", result[0].Name)
		fmt.Println("Grade\t:", result[0].Grade)
	}
}

func update() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	selector := bson.M{"name": "Wick"}
	changes := student{"John Wick", 2}

	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("update success!")
}

func remove() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	selector := bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("remove success!")
}

func aggregate() {
	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
	  [
			{
				"$group": {
					"_id": null,
					"Total": {
						"$sum": 1
					}
				}
			},
			{
				"$project": {
					"Total": 1,
					"_id": 0
				}
			}
		]
	`)), true, &pipeline)
	if err != nil {
		panic(err.Error())
	}

	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		panic(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total\t:", result[0]["Total"])
	}
}
