package repository

import (
	"github.com/aeramu/qiup-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type model struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	Avatar       string
	Body         string
	UpvoterIDs   map[string]bool    `bson:"upvoterIDs"`
	DownvoterIDs map[string]bool    `bson:"downvoterIDs"`
	ReplyCount   int                `bson:"replyCount"`
	ParentID     primitive.ObjectID `bson:"parentID"`
	RepostID     primitive.ObjectID `bson:"repostID"`
	RoomID       primitive.ObjectID `bson:"roomID"`
}

func newModel(name string, avatar string, body string, parentID string, repostID string, roomID string) *model {
	parentid, _ := primitive.ObjectIDFromHex(parentID)
	repostid, _ := primitive.ObjectIDFromHex(repostID)
	roomid, _ := primitive.ObjectIDFromHex(roomID)
	return &model{
		ID:           primitive.NewObjectID(),
		Name:         name,
		Avatar:       avatar,
		Body:         body,
		UpvoterIDs:   map[string]bool{},
		DownvoterIDs: map[string]bool{},
		ReplyCount:   0,
		ParentID:     parentid,
		RepostID:     repostid,
		RoomID:       roomid,
	}
}

func modelFromEntity(e entity.MenfessPost) *model {
	id, _ := primitive.ObjectIDFromHex(e.ID())
	parentID, _ := primitive.ObjectIDFromHex(e.ParentID())
	repostID, _ := primitive.ObjectIDFromHex(e.RepostID())
	roomID, _ := primitive.ObjectIDFromHex(e.RoomID())
	return &model{
		ID:           id,
		Name:         e.Name(),
		Avatar:       e.Avatar(),
		Body:         e.Body(),
		UpvoterIDs:   e.UpvoterIDs(),
		DownvoterIDs: e.DownvoterIDs(),
		ReplyCount:   e.ReplyCount(),
		ParentID:     parentID,
		RepostID:     repostID,
		RoomID:       roomID,
	}
}

func (m *model) Entity() entity.MenfessPost {
	return entity.MenfessPostConstructor{
		ID:           m.ID.Hex(),
		Timestamp:    int(m.ID.Timestamp().Unix()),
		Name:         m.Name,
		Avatar:       m.Avatar,
		Body:         m.Body,
		UpvoterIDs:   m.UpvoterIDs,
		DownvoterIDs: m.DownvoterIDs,
		ReplyCount:   m.ReplyCount,
		ParentID:     m.ParentID.Hex(),
		RepostID:     m.RepostID.Hex(),
		RoomID:       m.RoomID.Hex(),
	}.New()
}

func modelListToEntity(modelList []*model) []entity.MenfessPost {
	var entityList []entity.MenfessPost
	for _, model := range modelList {
		entityList = append(entityList, model.Entity())
	}
	return entityList
}

func idListFromHex(hexList []string) []primitive.ObjectID {
	var idList []primitive.ObjectID
	for _, hex := range hexList {
		id, _ := primitive.ObjectIDFromHex(hex)
		idList = append(idList, id)
	}
	return idList
}
