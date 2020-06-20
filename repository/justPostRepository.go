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
		database: client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type JustPostRepositoryImplementation struct{
	client *mongo.Client
	database *mongo.Database
	collection *mongo.Collection
}

func (repository *JustPostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{})(*entity.JustPost){
	filter := bson.D{
		{indexName,indexValue},
	}
	var post entity.JustPost
	repository.collection.FindOne(context.TODO(), filter).Decode(&post)
	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *JustPostRepositoryImplementation) GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID)([]*entity.JustPost){
	filter := bson.D{
		{"$and", bson.A{
			bson.D{
				{indexName, indexValue},
			},
			bson.D{
				{"_id", bson.D{
					{"$gt", after},
				}},
			},
		}},
	}
	option := options.Find().SetLimit(int64(limit))
	cursor,_ := repository.collection.Find(context.TODO(), filter, option)

	var postList []*entity.JustPost
	cursor.All(context.TODO(), &postList)
	return postList
}

func (repository *JustPostRepositoryImplementation) GetDataList(limit int32,after primitive.ObjectID)([]*entity.JustPost){
	filter := bson.D{
		{"_id",bson.D{
			{"$lt",after},
		}},
	}
	sort:= bson.D{
		{"_id",-1},
	}
	option := options.Find().SetLimit(int64(limit)).SetSort(sort)
	cursor,_ := repository.collection.Find(context.TODO(), filter, option)
	
	var postList []*entity.JustPost
	cursor.All(context.TODO(), &postList)
	return postList
}

func (repository *JustPostRepositoryImplementation) PutData(post *entity.JustPost){
	filter := bson.D{
		{"_id", post.ParentID},
	}
	update := bson.D{
		{"$inc",bson.D{
			{"replyCount",1},
		}},
	}
	option := options.BulkWrite().SetOrdered(false)
	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(post),
		mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true),
	}
	repository.collection.BulkWrite(context.TODO(), models, option)
}