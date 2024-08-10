package test

import (
	"context"
	"github.com/EnesHarman/eventify/internal/model"
	"github.com/EnesHarman/eventify/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

var eventRepositoryTest = repository.NewEventRepository()

//func TestMain(m *testing.M) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	clientOptions := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		panic("Can not connect to the mongo!")
//	}
//	fmt.Println(client)
//	exitCode := m.Run()
//	disconnect(client, ctx)
//	os.Exit(exitCode)
//}

func TestInsertOne(t *testing.T) {
	event := model.Event{
		Code:   "n:op",
		UserId: "enesharman2",
		Ts:     time.Now(),
	}
	t.Run("InsertEvent", func(t *testing.T) {
		err := eventRepositoryTest.InsertEvent(event)
		assert.Nil(t, err)
	})
}

func TestGetEvents(t *testing.T) {
	page, size := int64(1), int64(10)
	t.Run("GetEvents", func(t *testing.T) {
		events, err := eventRepositoryTest.GetEvents(page, size)
		assert.Nil(t, err)
		assert.NotNil(t, events)
	})
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
