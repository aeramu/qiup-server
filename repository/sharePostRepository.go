package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type SharePostRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.SharePost)
	PutData(account *entity.SharePost)
	//UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*entity.ShareAccount)
}

func NewSharePostRepository()(SharePostRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &SharePostRepositoryImplementation{
		client: client,
	}
}

type SharePostRepositoryImplementation struct{
	client *mongo.Client
}

func (repository *SharePostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) (*entity.SharePost){
	collection := repository.client.Database("qiup").Collection("sharePost")
	
	var post entity.SharePost
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&post)

	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *SharePostRepositoryImplementation) PutData(post *entity.SharePost){
	collection := repository.client.Database("qiup").Collection("sharePost")
	collection.InsertOne(context.TODO(),post)
}