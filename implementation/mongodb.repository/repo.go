package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

//New MenfessPostRepo Constructor
func New() usecase.MenfessPostRepo {
	if client == nil {
		client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
			"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
		))
	}
	return &menfessPostRepo{
		client:     client,
		database:   client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type menfessPostRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (repo *menfessPostRepo) GetDataByID(id string) *entity.MenfessPost {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", objectID},
	}
	var post entity.MenfessPost
	repo.collection.FindOne(context.TODO(), filter).Decode(&post)
	if post.ID.IsZero() {
		return nil
	}
	return &post
}
