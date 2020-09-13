package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aeramu/qiup-server/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

//New MenfessPostRepo Constructor
func New() usecase.Repository {
	if client == nil {
		client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
			"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
		))
	}
	return &menfessRepo{
		client:     client,
		database:   client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type menfessRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (repo *menfessRepo) NewID() string {
	return primitive.NewObjectID().Hex()
}