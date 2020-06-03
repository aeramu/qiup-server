package repository

import (
	"context"
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileRepository interface{
	GetDataByIndex(indexName string, indexValue interface{}) (*entity.Profile)
	PutData(profile *entity.Profile)
	UpdateData(profileID string, indexName string, indexValue interface{}) (*entity.Profile)
}

func NewProfileRepository()(ProfileRepository){
	client,_ := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &ProfileRepositoryImplementation{
		client: client,
	}
}

type ProfileRepositoryImplementation struct{
	client *mongo.Client
}

func (repository *ProfileRepositoryImplementation) GetDataByIndex(indexName string, indexValue interface{}) (*entity.Profile){
	collection := repository.client.Database("qiup").Collection("profile")
	
	var profile entity.Profile
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&profile)

	if profile.ID == "" {
		return nil
	}
	return &profile
}

func (repository *ProfileRepositoryImplementation) PutData(profile *entity.Profile){
	collection := repository.client.Database("qiup").Collection("profile")
	collection.InsertOne(context.TODO(),profile)
}

func (repository *ProfileRepositoryImplementation) UpdateData(profileID string, indexName string, indexValue interface{}) (*entity.Profile){
	collection := repository.client.Database("qiup").Collection("profile")
	collection.UpdateOne(context.TODO(),bson.D{{"_id",profileID}},bson.D{{"$set", bson.D{{indexName, indexValue}}}})

	var profile entity.Profile
	collection.FindOne(context.TODO(),bson.D{{indexName,indexValue}}).Decode(&profile)

	return &profile
}