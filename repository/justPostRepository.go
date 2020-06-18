package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JustPostRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.JustPost)
	GetDataList(limit int32, after primitive.ObjectID) ([]*entity.JustPost)
	GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID) ([]*entity.JustPost)
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

func (repository *JustPostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{})(*entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")

	filter := bson.D{{indexName,indexValue}}
	
	var post entity.JustPost
	collection.FindOne(context.TODO(), filter).Decode(&post)

	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *JustPostRepositoryImplementation) GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID)([]*entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")

	afterFilter := bson.D{{"_id",bson.D{{"$gt",after}}}}
	indexFilter := bson.D{{indexName,indexValue}}
	filter := bson.D{{"$and",bson.A{indexFilter,afterFilter}}}

	option := options.Find().SetLimit(int64(limit))

	cursor,_ := collection.Find(context.TODO(), filter, option)

	var postList []*entity.JustPost
	cursor.All(context.TODO(), &postList)

	return postList
}

func (repository *JustPostRepositoryImplementation) GetDataList(limit int32,after primitive.ObjectID)([]*entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")

	filter := bson.D{{"_id",bson.D{{"$lt",after}}}}

	option := options.Find()
	option.SetLimit(int64(limit))
	option.SetSort(bson.D{{"_id",-1}})

	cursor,_ := collection.Find(context.TODO(), filter, option)
	
	var postList []*entity.JustPost
	cursor.All(context.TODO(), &postList)

	return postList
}

func (repository *JustPostRepositoryImplementation) PutData(post *entity.JustPost){
	collection := repository.client.Database("qiup").Collection("justPost")
	collection.InsertOne(context.TODO(),post)
}