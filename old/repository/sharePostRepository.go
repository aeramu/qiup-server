package repository

import (
	"context"

	"github.com/aeramu/qiup-server/old/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//SharePostRepository interface
type SharePostRepository interface {
	GetDataByIndex(indexName string, indexValue interface{}) *domain.SharePost
	PutData(account *domain.SharePost)
	GetDataList() []*domain.SharePost
	//UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*domain.ShareAccount)
}

//NewSharePostRepository constructor
func NewSharePostRepository() SharePostRepository {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &sharePostRepositoryImplementation{
		client: client,
	}
}

type sharePostRepositoryImplementation struct {
	client *mongo.Client
}

func (repository *sharePostRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) *domain.SharePost {
	collection := repository.client.Database("qiup").Collection("sharePost")

	var post domain.SharePost
	collection.FindOne(context.TODO(), bson.D{{indexName, indexValue}}).Decode(&post)

	if post.ID.IsZero() {
		return nil
	}
	return &post
}

func (repository *sharePostRepositoryImplementation) PutData(post *domain.SharePost) {
	collection := repository.client.Database("qiup").Collection("sharePost")
	collection.InsertOne(context.TODO(), post)
}

func (repository *sharePostRepositoryImplementation) GetDataList() []*domain.SharePost {
	collection := repository.client.Database("qiup").Collection("sharePost")

	var postList []*domain.SharePost
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	cursor.All(context.TODO(), &postList)

	return postList
}
