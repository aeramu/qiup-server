package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ID is function to convert from Hex to MongoDB ObjectID
func ID(hex string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(hex)
	return id
}

//NewID is function to create new MongoDB ObjectID
func NewID() primitive.ObjectID {
	return primitive.NewObjectID()
}

//Account is db model of qiup account in repository
type Account struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string
	Password string
}

//ShareAccount is db model of share account in repository
type ShareAccount struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string
	ShareProfile ShareProfile `bson:"shareProfile"`
}

//ShareProfile is db model of share profile in repository
type ShareProfile struct {
	Name         string
	Bio          string
	ProfilePhoto string
	CoverPhoto   string
}

//SharePost is db model of share post in repository
type SharePost struct {
	ID        primitive.ObjectID `bson:"_id"`
	AccountID primitive.ObjectID
	Body      string
}
