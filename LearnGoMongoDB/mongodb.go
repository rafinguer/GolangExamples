package main

// Reference: https://labix.org/mgo
// Documentation: https://pkg.go.dev/gopkg.in/mgo.v2

// Reference: https://www.mongodb.com/docs/drivers/go/current/quick-start/#std-label-golang-quickstart

// Official MongoDB Driver installation:
// 1) Access to $GO_PATH/src
// 2) $ go get go.mongodb.org/mongo-driver/mongo
// 3) $ go get github.com/joho/godotenv   // Dependencies
// If inconsistent errors has occurred (inconsistent vendoring and explicity required in go.mod), fix it with:
// $ go mod vendor

// Mongo Driver installation
// 1) Access to $GO_PATH/src
// 2) $ go get gopkg.in/mgo.v2
// If inconsistent errors has occurred (inconsistent vendoring and explicity required in go.mod), fix it with:
// $ go mod vendor

// For this example, MongoDB is available via Docker
// $ docker pull mongo
//  --- If admin user defined --- $ docker run --name mongodb -e MONGO_INITDB_DATABASE=example-db -e MONGO_INITDB_ROOT_USERNAME=mongo-root -e MONGO_INITDB_ROOT_PASSWORD=r00t-p4ssw0rd -p 27017:27017 -d mongo
// $ docker run --name mongodb -p 27017:27017 -d mongo
// Note: If you want to know the ip used by the container: $ docker inspect mongodb | grep "IPAdress"
// $ docker exec -it mongodb /bin/bash
// --- If admin yuser defined --- $ mongosh -u mongo-root -p r00t-p4ssw0rd
// $ mongosh
// > show dbs
// > use example-db
// > db.createCollection("users")
// > show collections
// > db.users.insertOne({"firstname":"Rob", "lastname": "Pike", "username": "rob"})
// > db.users.insertOne({"firstname":"Ken", "lastname": "Thompson", "username": "ken"})
// > db.users.insertOne({"firstname":"Robert", "lastname": "Griesemet", "username": "gri"})
// > db.users.find({})

import (

	//	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	//"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/example-db"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("example-db").Collection("users")
	fmt.Println("Type: ", reflect.TypeOf(coll))

	user := findUserBy(coll, "rob")
	fmt.Println("\n---\nUser found: ", user)

	users := getUsers(coll)
	for _, i := range users {
		fmt.Println("\n---\nUser: ", i)
	}

	newUser := bson.D{{"firstname", "Rafael"}, {"lastname", "Hernamperez"}, {"username", "raf"}}
	isUserAdded := addUser(coll, &newUser)

	if !isUserAdded {
		fmt.Println("\n---\nNew user could not be added")
	} else {
		fmt.Println("\n---\nNew user add: ", isUserAdded)
	}

	users = getUsers(coll)
	for _, i := range users {
		fmt.Println("\n---\nUser: ", i)
	}

}

// This function finds a given user by its username
func findUserBy(coll *mongo.Collection, username string) string {
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the username %s\n", username)
		return fmt.Sprintf("No document was found with the username %s\n", username)
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", jsonData)
	return string(jsonData)
}

// This function returns all the users in 'users' collection
// To achieve this, the filter must be {} (everything)
func getUsers(coll *mongo.Collection) []string {
	var result []string
	var rows []bson.M
	filter := bson.D{} // Filter
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), &rows); err != nil {
		panic(err)
	}

	for _, row := range rows {
		jsonRow, err := json.MarshalIndent(row, "", " ")
		if err != nil {
			panic(err)
		}
		result = append(result, string(jsonRow))
	}

	return result
}

// This function returns adds a new user in 'users' collection
func addUser(coll *mongo.Collection, user *bson.D) bool {
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n---\nID User added: ", result.InsertedID)

	return true
}
