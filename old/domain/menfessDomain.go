package domain

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

// GetID node interface
func (mp *MenfessPost) GetID() primitive.ObjectID {
	return mp.ID
}

// NewMenfessPost constructor
func NewMenfessPost(accountID primitive.ObjectID) *MenfessPost {
	return &MenfessPost{
		ID:        NewID(),
		AccountID: accountID,
	}
}

// SetName setter name
func (mp *MenfessPost) SetName(name string) *MenfessPost {
	mp.Name = name
	return mp
}

// SetAvatar setter avatar
func (mp *MenfessPost) SetAvatar(avatar string) *MenfessPost {
	mp.Avatar = avatar
	return mp
}

// SetBody setter body
func (mp *MenfessPost) SetBody(body string) *MenfessPost {
	mp.Body = body
	return mp
}

// SetParentID setter parentID
func (mp *MenfessPost) SetParentID(parentID primitive.ObjectID) *MenfessPost {
	mp.ParentID = parentID
	return mp
}

// Timestamp get timestamp from id
func (mp *MenfessPost) Timestamp() int64 {
	return mp.ID.Timestamp().Unix()
}
