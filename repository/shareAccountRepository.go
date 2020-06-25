package repository

import (
	"context"

	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ShareAccountRepository interface
type ShareAccountRepository interface {
	GetDataByIndex(indexName string, indexValue interface{}) *entity.ShareAccount
	PutData(account *entity.ShareAccount)
	UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) *entity.ShareAccount
}

//NewShareAccountRepository constructor
func NewShareAccountRepository() ShareAccountRepository {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &shareAccountRepositoryImplementation{
		client: client,
	}
}

type shareAccountRepositoryImplementation struct {
	client *mongo.Client
}

func (repository *shareAccountRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) *entity.ShareAccount {
	collection := repository.client.Database("qiup").Collection("shareAccount")

	var account entity.ShareAccount
	collection.FindOne(context.TODO(), bson.D{{indexName, indexValue}}).Decode(&account)

	if account.ID.IsZero() {
		return nil
	}
	return &account
}

func (repository *shareAccountRepositoryImplementation) PutData(account *entity.ShareAccount) {
	collection := repository.client.Database("qiup").Collection("shareAccount")
	collection.InsertOne(context.TODO(), account)
}

func (repository *shareAccountRepositoryImplementation) UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) *entity.ShareAccount {
	collection := repository.client.Database("qiup").Collection("shareAccount")
	collection.UpdateOne(context.TODO(), bson.D{{"_id", accountID}}, bson.D{{"$set", bson.D{{indexName, indexValue}}}})

	var account entity.ShareAccount
	collection.FindOne(context.TODO(), bson.D{{"_id", accountID}}).Decode(&account)

	return &account
}
