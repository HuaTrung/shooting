package database

import (
	"context"
	"fmt"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)
// Conn represent a SQL connection.
type ConnMongo struct {
	db *mongo.Client
}


// New creates a new SQL connection.
func NewMongo( uri string) (*mongo.Client, error) {
	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return client, nil
}

// Close ensures that the connection is terminated.
func (c *ConnMongo) Close() error {
	var err error
	if err = c.db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	return err
}