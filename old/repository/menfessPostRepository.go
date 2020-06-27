package repository

import (
	"context"

	"github.com/aeramu/qiup-server/old/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MenfessPostRepository interface
type MenfessPostRepository interface {
	GetDataByIndex(indexName string, indexValue interface{}) *domain.MenfessPost
	GetDataList(limit int32, after primitive.ObjectID) []*domain.MenfessPost
	GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID, sort int32) []*domain.MenfessPost
	PutData(account *domain.MenfessPost)
	//UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*domain.ShareAccount)
}

var client *mongo.Client = nil

// NewMenfessPostRepository constructor with singleton client
func NewMenfessPostRepository() MenfessPostRepository {
	if client == nil {
		client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
			"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
		))
	}
	return &MenfessPostRepositoryImplementation{
		client:     client,
		database:   client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

// MenfessPostRepositoryImplementation implement
type MenfessPostRepositoryImplementation struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// GetDataByIndex Get
func (repository *MenfessPostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) *domain.MenfessPost {
	filter := bson.D{
		{indexName, indexValue},
	}
	var post domain.MenfessPost
	repository.collection.FindOne(context.TODO(), filter).Decode(&post)
	if post.ID.IsZero() {
		return nil
	}
	return &post
}

// GetDataListByIndex get
func (repository *MenfessPostRepositoryImplementation) GetDataListByIndex(indexName string, indexValue interface{}, limit int32, after primitive.ObjectID, sort int32) []*domain.MenfessPost {
	comparator := "$gt"
	if sort == -1 {
		comparator = "$lt"
	}
	filter := bson.D{
		{"$and", bson.A{
			bson.D{
				{indexName, indexValue},
			},
			bson.D{
				{"_id", bson.D{
					{comparator, after},
				}},
			},
		}},
	}
	sortOpt := bson.D{
		{"_id", sort},
	}
	option := options.Find().SetLimit(int64(limit)).SetSort(sortOpt)
	cursor, _ := repository.collection.Find(context.TODO(), filter, option)

	var postList []*domain.MenfessPost
	cursor.All(context.TODO(), &postList)
	return postList
}

// GetDataList get
func (repository *MenfessPostRepositoryImplementation) GetDataList(limit int32, after primitive.ObjectID) []*domain.MenfessPost {
	filter := bson.D{
		{"_id", bson.D{
			{"$lt", after},
		}},
	}
	sort := bson.D{
		{"_id", -1},
	}
	option := options.Find().SetLimit(int64(limit)).SetSort(sort)
	cursor, _ := repository.collection.Find(context.TODO(), filter, option)

	var postList []*domain.MenfessPost
	cursor.All(context.TODO(), &postList)
	return postList
}

//PutData put
func (repository *MenfessPostRepositoryImplementation) PutData(post *domain.MenfessPost) {
	filter := bson.D{
		{"_id", post.ParentID},
	}
	update := bson.D{
		{"$inc", bson.D{
			{"replyCount", 1},
		}},
	}
	option := options.BulkWrite().SetOrdered(false)
	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(post),
		mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true),
	}
	repository.collection.BulkWrite(context.TODO(), models, option)
}
