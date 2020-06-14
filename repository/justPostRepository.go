package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type JustPostRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.JustPost)
	GetDataListByIndex(indexName string, indexValue interface{}, limit int32) ([]*entity.JustPost)
	PutData(account *entity.JustPost)
	//UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*entity.ShareAccount)
}

func NewJustPostRepository()(JustPostRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &JustPostRepositoryImplementation{
		client: client,
	}
}

type JustPostRepositoryImplementation struct{
	client *mongo.Client
}

func (repository *JustPostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) (*entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")
	
	var post entity.JustPost
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&post)

	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *JustPostRepositoryImplementation) GetDataListByIndex(indexName string, indexValue interface{}, limit int32) ([]*entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")

	var postList []*entity.JustPost
	cursor,_ := collection.Find(context.TODO(),bson.D{{indexName,indexValue}},options.Find().SetLimit(int64(limit)))
	cursor.All(context.TODO(),&postList)

	return postList
}

func (repository *JustPostRepositoryImplementation) PutData(post *entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")
	collection.InsertOne(context.TODO(),post)
}