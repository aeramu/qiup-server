package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenfessPostRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.MenfessPost)
	GetDataList(limit int32, after primitive.ObjectID) ([]*entity.MenfessPost)
	GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID) ([]*entity.MenfessPost)
	PutData(account *entity.MenfessPost)
	//UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*entity.ShareAccount)
}

func NewMenfessPostRepository()(MenfessPostRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &MenfessPostRepositoryImplementation{
		client: client,
		database: client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type MenfessPostRepositoryImplementation struct{
	client *mongo.Client
	database *mongo.Database
	collection *mongo.Collection
}

func (repository *MenfessPostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{})(*entity.MenfessPost){
	filter := bson.D{
		{indexName,indexValue},
	}
	var post entity.MenfessPost
	repository.collection.FindOne(context.TODO(), filter).Decode(&post)
	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *MenfessPostRepositoryImplementation) GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID)([]*entity.MenfessPost){
	filter := bson.D{
		{"$and", bson.A{
			bson.D{
				{indexName, indexValue},
			},
			bson.D{
				{"_id", bson.D{
					{"$lt", after},
				}},
			},
		}},
	}
	sort:= bson.D{
		{"_id",-1},
	}
	option := options.Find().SetLimit(int64(limit)).SetSort(sort)
	cursor,_ := repository.collection.Find(context.TODO(), filter, option)

	var postList []*entity.MenfessPost
	cursor.All(context.TODO(), &postList)
	return postList
}

func (repository *MenfessPostRepositoryImplementation) GetDataList(limit int32,after primitive.ObjectID)([]*entity.MenfessPost){
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
	
	var postList []*entity.MenfessPost
	cursor.All(context.TODO(), &postList)
	return postList
}

func (repository *MenfessPostRepositoryImplementation) PutData(post *entity.MenfessPost){
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