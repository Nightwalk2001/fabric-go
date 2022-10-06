package mongo

import (
	"context"
	"time"

	"fabric/config"
	gomongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client  *gomongo.Client
	Clothes *gomongo.Collection
	Clients *gomongo.Collection
	Bills   *gomongo.Collection
)

func Setup(conf *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(conf.Uri)
	client, _ = gomongo.Connect(ctx, opts)
	db := client.Database("fabric")

	Clothes = db.Collection("clothes")
	Clients = db.Collection("clients")
	Bills = db.Collection("bills")
}

func Disconnect() {
	_ = client.Disconnect(context.Background())
	Clothes = nil
	Clients = nil
	Bills = nil
}
