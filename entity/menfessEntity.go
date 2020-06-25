package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MenfessPost is db model of menfess post in repository
type MenfessPost struct {
	ID            primitive.ObjectID `bson:"_id"`
	AccountID     primitive.ObjectID `bson:"accountID"`
	Name          string
	Avatar        string
	Body          string
	ReplyCount    int32              `bson:"replyCount"`
	UpvoteCount   int32              `bson:"upvoteCount"`
	DownvoteCount int32              `bson:"downvoteCount"`
	ParentID      primitive.ObjectID `bson:"parentID"`
}

//NewMenfessPost constructor
func NewMenfessPost(name string, avatar string, body string) *MenfessPost {
	return &MenfessPost{
		ID:     NewID(),
		Name:   name,
		Avatar: avatar,
		Body:   body,
	}
}

//SetParentID Setter parentID
func (mp *MenfessPost) SetParentID(parentID primitive.ObjectID) *MenfessPost {
	mp.ParentID = parentID
	return mp
}
