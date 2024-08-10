package repository

import (
	"context"
	"fmt"
	"github.com/EnesHarman/eventify/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type EventRepository interface {
	InsertEvent(event model.Event) error
	GetEvents(page, size int64) ([]model.Event, error)
}

type MongoEventRepository struct {
}

func (repository MongoEventRepository) InsertEvent(event model.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer disconnect(client, ctx)

	collection := client.Database("event").Collection("events")
	_, err = collection.InsertOne(ctx, bson.D{{"code", event.Code}, {"userId", event.UserId}, {"ts", event.Ts}})
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}

	return nil
}

func (repository MongoEventRepository) GetEvents(page, size int64) ([]model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer disconnect(client, ctx)

	collection := client.Database("event").Collection("events")
	cursor, err := collection.Find(ctx, bson.D{}, options.Find().SetSkip((page-1)*size).SetLimit(size))
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	events := make([]model.Event, 0)
	if err = cursor.All(ctx, &events); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %v", err)
	}

	return events, nil

}

func NewEventRepository() EventRepository {
	return &MongoEventRepository{}
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
