package DBConnector

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var mongoClient *mongo.Client

func LoadEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file!")
	} else {
		fmt.Println(" .env file loaded!")
		userName := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
		password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
		uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", userName, password)
		return uri
	}
	return ""
}
func ConnectToMongo() error {
	const uri = "mongodb://localhost:27017"
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(LoadEnv()).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	mongoClient = client
	fmt.Println("Connected to mongo!")
	return nil
}

func DisconnectFromMongo() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		fmt.Println("Failed to disconnect from mongo! Error:", err)
		panic(err)
	}
	fmt.Println("Disconnected from mongo!")
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}
