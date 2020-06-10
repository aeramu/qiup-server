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