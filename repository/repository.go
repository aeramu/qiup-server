package repository

import (
	"context"

	"gitlab.com/kentanggoreng/quip-server/entity"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// interface for account repository
type AccountRepository interface{
	GetDataByIndex(indexName string, indexValue string) (*entity.Account, error)
}

// Constructor for AccountRepository
func NewAccountRepository()(AccountRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@quip-wrbox.mongodb.net/",
	))
	return &AccountRepositoryImplementation{
		client: client,
	}
}

// Class for account repository implementation
type AccountRepositoryImplementation struct{
	client *mongo.Client
}

func (repository *AccountRepositoryImplementation) GetDataByIndex(indexName string, indexValue string) (*entity.Account, error){
	collection := repository.client.Database("quip").Collection("account")
	
	var account entity.Account
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&account)

	if account.ID == "" {
		return nil, nil
	}
	return &account, nil
}