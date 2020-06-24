package entity

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenfessPost struct{
	ID primitive.ObjectID `bson:"_id"`
	AccountID primitive.ObjectID `bson:"accountID"`
	Name string
	Avatar string
	Body string
	ReplyCount int32 `bson:"replyCount"`
	UpvoteCount int32 `bson:"upvoteCount"`
	DownvoteCount int32 `bson:"downvoteCount"`
	ParentID primitive.ObjectID `bson:"parentID"`
}