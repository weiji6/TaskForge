package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient holds the client connection to MongoDB.
type MongoDBClient struct {
	Client *mongo.Client
}

// NewMongoDBClient initializes a new MongoDB client.
func NewMongoDBClient(uri string) (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &MongoDBClient{Client: client}, nil
}

// Message represents a message in the repository.
type Message struct {
	ID      string `bson:"_id,omitempty"`
	Content string `bson:"content"`
}

// MessageRepository defines the interface for message operations.
type MessageRepository interface {
	Save(message Message) error
	FindAll() ([]Message, error)
}

// MongoDBMessageRepository implements MessageRepository for MongoDB.
type MongoDBMessageRepository struct {
	client *MongoDBClient
}

// NewMongoDBMessageRepository creates a new MongoDBMessageRepository.
func NewMongoDBMessageRepository(client *MongoDBClient) *MongoDBMessageRepository {
	return &MongoDBMessageRepository{client: client}
}

// Save saves a message to the MongoDB repository.
func (repo *MongoDBMessageRepository) Save(message Message) error {
	collection := repo.client.Client.Database("example_db").Collection("messages")
	_, err := collection.InsertOne(context.Background(), message)
	return err
}

// FindAll retrieves all messages from the MongoDB repository.
func (repo *MongoDBMessageRepository) FindAll() ([]Message, error) {
	collection := repo.client.Client.Database("example_db").Collection("messages")
	var messages []Message
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var message Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}