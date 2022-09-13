package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClientConversation(url string) *MongoConnection {
	var err error
	var ctx context.Context

	mc := &MongoConnection{
		database:   "telegom",
		collection: "conversations",
	}

	mc.clientOptions = options.Client().ApplyURI(url)

	mc.clientContext, mc.clientCancelContext = context.WithCancel(context.Background())

	mc.client, err = mongo.Connect(ctx, mc.clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	return mc
}

type MongoConnection struct {
	client              *mongo.Client
	clientOptions       *options.ClientOptions
	clientCancelContext context.CancelFunc
	clientContext       context.Context
	database            string
	collection          string
}

func (mc *MongoConnection) CancelConection() {
	mc.clientCancelContext()
}

func (mc *MongoConnection) FindConversation(chat_id int) (d *CommandsPending, err error) {
	var result bson.M

	collect := *mc.client.Database(mc.database).Collection(mc.collection)

	err = collect.FindOne(mc.clientContext, bson.D{{Key: "chat_id", Value: chat_id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	bsonData, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(bsonData, &d)
	if err != nil {
		return nil, err
	}

	return
}
