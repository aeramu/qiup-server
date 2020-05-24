package repository

import (
	"gitlab.com/kentanggoreng/quip-server/entity"
  
	"context"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// interface for account repository
type AccountRepository interface{
	GetDataByIndex(indexName string, indexValue string) (*entity.Account, error)
}

// Constructor for AccountRepository
func NewAccountRepository()(AccountRepository){
	credential := option.WithCredentialsFile("credential.json")
	app,_ := firebase.NewApp(context.Background(), nil, credential)
	client,_ := app.Firestore(context.Background())
	return &AccountRepositoryImplementation{
		client: client,
	}
}

// Class for account repository implementation
type AccountRepositoryImplementation struct{
	client *firestore.Client
}

func (repository *AccountRepositoryImplementation) GetDataByIndex(indexName string, indexValue string) (*entity.Account, error){
	data, err := repository.client.Collection("account").Where(indexName, "==", indexValue).Documents(context.Background()).Next()
	if err != nil{
		return nil, err
	}
	return &entity.Account{
		ID: data.Ref.ID,
		Email: data.Data()["Email"].(string),
		Username: data.Data()["Username"].(string),
		Password: data.Data()["Password"].(string),
	}, nil
}

