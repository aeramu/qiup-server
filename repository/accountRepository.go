package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// interface for account repository
type AccountRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.Account)
	PutData(account *entity.Account)
	UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*entity.Account)
}

// Constructor for AccountRepository
func NewAccountRepository()(AccountRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &AccountRepositoryImplementation{
		client: client,
	}
}

// Class for account repository implementation
type AccountRepositoryImplementation struct{
	client *mongo.Client
}

func (repository *AccountRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) (*entity.Account){
	collection := repository.client.Database("qiup").Collection("account")
	
	var account entity.Account
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&account)

	if account.ID.IsZero() {
		return nil
	}
	return &account
}

func (repository *AccountRepositoryImplementation) PutData(account *entity.Account){
	collection := repository.client.Database("qiup").Collection("account")
	collection.InsertOne(context.TODO(),account)
}

func (repository *AccountRepositoryImplementation) UpdateData(accountID primitive.ObjectID, indexName string, indexValue interface{}) (*entity.Account){
	collection := repository.client.Database("qiup").Collection("account")
	collection.UpdateOne(context.TODO(),bson.D{{"_id",accountID}},bson.D{{"$set", bson.D{{indexName, indexValue}}}})

	var account entity.Account
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&account)

	return &account
}