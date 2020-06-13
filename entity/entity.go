package entity

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct{
	ID primitive.ObjectID `bson:"_id"`
	Email string
	Password string
}

type ShareAccount struct{
	ID primitive.ObjectID `bson:"_id"`
	Username string
	ShareProfile ShareProfile `bson:"shareProfile"`
}

type ShareProfile struct{
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
}

type SharePost struct{
	ID primitive.ObjectID `bson:"_id"`
	AccountID primitive.ObjectID
	Body string
}

type JustPost struct{
	ID primitive.ObjectID `bson:"_id"`
	ParentID primitive.ObjectID
	Name string
	Body string
}