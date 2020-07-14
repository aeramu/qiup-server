package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

//New MenfessPostRepo Constructor
func New() usecase.MenfessRepo {
	if client == nil {
		client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
			"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
		))
	}
	return &menfessRepo{
		client:     client,
		database:   client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type menfessRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (repo *menfessRepo) NewID() string {
	return primitive.NewObjectID().Hex()
}

func (repo *menfessRepo) GetPostByID(id string) entity.MenfessPost {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objectID}}
	var post post
	repo.collection.FindOne(context.TODO(), filter).Decode(&post)

	if post.ID.IsZero() {
		return nil
	}
	return post.Entity()
}

func (repo *menfessRepo) GetPostListByParentID(parentID string, first int, after string, ascSort bool) []entity.MenfessPost {
	parentid, _ := primitive.ObjectIDFromHex(parentID)
	afterid, _ := primitive.ObjectIDFromHex(after)
	comparator := "$lt"
	sort := -1
	if ascSort {
		comparator = "$gt"
		sort = 1
	}

	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"parentID", parentid}},
			bson.D{
				{"_id", bson.D{
					{comparator, afterid},
				}},
			},
		}},
	}
	sortOpt := bson.D{{"_id", sort}}
	option := options.Find().SetLimit(int64(first)).SetSort(sortOpt)
	cursor, _ := repo.collection.Find(context.TODO(), filter, option)

	var postList []*post
	cursor.All(context.TODO(), &postList)
	return modelListToEntity(postList)
}

func (repo *menfessRepo) GetPostListByRoomIDs(roomIDs []string, first int, after string, ascSort bool) []entity.MenfessPost {
	roomids := idListFromHex(roomIDs)
	afterid, _ := primitive.ObjectIDFromHex(after)
	comparator := "$lt"
	sort := -1
	if ascSort {
		comparator = "$gt"
		sort = 1
	}

	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"roomID", bson.D{
				{"$in", roomids},
			}}},
			bson.D{{"_id", bson.D{
				{comparator, afterid},
			}}},
		}},
	}
	sortOpt := bson.D{{"_id", sort}}
	option := options.Find().SetLimit(int64(first)).SetSort(sortOpt)
	cursor, _ := repo.collection.Find(context.TODO(), filter, option)

	var postList []*post
	cursor.All(context.TODO(), &postList)
	return modelListToEntity(postList)
}

func (repo *menfessRepo) PutPost(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost {
	model := newModel(name, avatar, body, parentID, repostID, roomID)
	filter := bson.D{{"_id", model.ParentID}}
	update := bson.D{
		{"$inc", bson.D{
			{"replyCount", 1},
		}},
	}
	option := options.BulkWrite().SetOrdered(false)
	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(model),
		mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true),
	}
	repo.collection.BulkWrite(context.TODO(), models, option)

	return model.Entity()
}

func (repo *menfessRepo) UpdateUpvoterIDs(postID string, accountID string, exist bool) {
	operator := "$set"
	if exist {
		operator = "$unset"
	}
	postid, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.D{{"_id", postid}}
	update := bson.D{
		{operator, bson.D{
			{"upvoterIDs." + accountID, true},
		}},
	}
	repo.collection.UpdateOne(context.TODO(), filter, update)
}

func (repo *menfessRepo) UpdateDownvoterIDs(postID string, accountID string, exist bool) {
	operator := "$set"
	if exist {
		operator = "$unset"
	}
	postid, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.D{{"_id", postid}}
	update := bson.D{
		{operator, bson.D{
			{"downvoterIDs." + accountID, true},
		}},
	}
	repo.collection.UpdateOne(context.TODO(), filter, update)
}

func (repo *menfessRepo) GetRoomList() []entity.MenfessRoom {
	filter := bson.D{{}}
	cursor, _ := repo.client.Database("menfess").Collection("room").Find(context.TODO(), filter)

	var modelList []*room
	cursor.All(context.TODO(), &modelList)
	return roomListToEntity(modelList)
}
