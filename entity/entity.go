package entity

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ID(hex string)(primitive.ObjectID){
	id,_ := primitive.ObjectIDFromHex(hex)
	return id
}
func NewID()(primitive.ObjectID){
	return primitive.NewObjectID()
}

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
	ParentID primitive.ObjectID `bson:"parentID"`
	Name string
	Avatar string
	Body string
}