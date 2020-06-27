package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

//New MenfessPostRepo Constructor
func New() usecase.MenfessPostRepo {
	if client == nil {
		client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
			"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
		))
	}
	return &menfessPostRepo{
		client:     client,
		database:   client.Database("qiup"),
		collection: client.Database("qiup").Collection("justPost"),
	}
}

type menfessPostRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (repo *menfessPostRepo) NewID() (string, int) {
	id := primitive.NewObjectID()
	return id.Hex(), int(id.Timestamp().Unix())
}

func (repo *menfessPostRepo) GetDataByID(id string) *entity.MenfessPost {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", objectID},
	}
	var post menfessPost
	repo.collection.FindOne(context.TODO(), filter).Decode(&post)
	if post.ID.IsZero() {
		return nil
	}
	return post.entity()
}

func (repo *menfessPostRepo) GetDataListByParentID(parentID string, first int, after string, ascSort bool) []*entity.MenfessPost {
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
			bson.D{
				{"parentID", parentid},
			},
			bson.D{
				{"_id", bson.D{
					{comparator, afterid},
				}},
			},
		}},
	}
	sortOpt := bson.D{
		{"_id", sort},
	}
	option := options.Find().SetLimit(int64(first)).SetSort(sortOpt)
	cursor, _ := repo.collection.Find(context.TODO(), filter, option)

	var postList []*menfessPost
	cursor.All(context.TODO(), &postList)
	fmt.Println(postList)

	return postListToEntity(postList)
}

func (repo *menfessPostRepo) PutData(e *entity.MenfessPost) {
	id, _ := primitive.ObjectIDFromHex(e.ID)
	parentID, _ := primitive.ObjectIDFromHex(e.ParentID)
	post := &menfessPost{
		ID:         id,
		Name:       e.Name,
		Avatar:     e.Avatar,
		Body:       e.Body,
		ReplyCount: e.ReplyCount,
		ParentID:   parentID,
	}
	filter := bson.D{
		{"_id", post.ParentID},
	}
	update := bson.D{
		{"$inc", bson.D{
			{"replyCount", 1},
		}},
	}
	option := options.BulkWrite().SetOrdered(false)
	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(post),
		mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true),
	}
	repo.collection.BulkWrite(context.TODO(), models, option)
}

func (repo *menfessPostRepo) Vote(postID string, accountID string, isUpvote bool) {
	listName := "downvoterIDs"
	if isUpvote {
		listName = "upvoterIDs"
	}
	accountid, _ := primitive.ObjectIDFromHex(accountID)
	postid, _ := primitive.ObjectIDFromHex(postID)
	filter := bson.D{
		{"_id", postid},
	}
	update := bson.D{
		{"$addToSet", bson.D{
			{listName, accountid},
		}},
	}
	repo.collection.UpdateOne(context.TODO(), filter, update)
}

func (repo *menfessPostRepo) Unvote(postID string, accountID string, isUpvote bool) {
	listName := "downvoterIDs"
	if isUpvote {
		listName = "upvoterIDs"
	}
	accountid, _ := primitive.ObjectIDFromHex(accountID)
	postid, _ := primitive.ObjectIDFromHex(postID)
	filter := bson.D{
		{"_id", postid},
	}
	update := bson.D{
		{"$pull", bson.D{
			{listName, accountid},
		}},
	}
	repo.collection.UpdateOne(context.TODO(), filter, update)
}

type menfessPost struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	Avatar       string
	Body         string
	UpvoterIDs   []primitive.ObjectID `bson:"upvoterIDs"`
	DownvoterIDs []primitive.ObjectID `bson:"downvoterIDs"`
	ReplyCount   int                  `bson:"replyCount"`
	ParentID     primitive.ObjectID   `bson:"parentID"`
}

func (m *menfessPost) entity() *entity.MenfessPost {
	return &entity.MenfessPost{
		ID:            m.ID.Hex(),
		Timestamp:     int(m.ID.Timestamp().Unix()),
		Name:          m.Name,
		Avatar:        m.Avatar,
		Body:          m.Body,
		UpvoterIDs:    idListToEntity(m.UpvoterIDs),
		DownvoterIDs:  idListToEntity(m.DownvoterIDs),
		UpvoteCount:   len(m.UpvoterIDs),
		DownvoteCount: len(m.DownvoterIDs),
		ReplyCount:    m.ReplyCount,
		ParentID:      m.ParentID.Hex(),
	}
}

func idListToEntity(list []primitive.ObjectID) []string {
	var idList []string
	for _, id := range list {
		idList = append(idList, id.Hex())
	}
	return idList
}

func postListToEntity(list []*menfessPost) []*entity.MenfessPost {
	var postList []*entity.MenfessPost
	for _, post := range list {
		postList = append(postList, post.entity())
	}
	return postList
}
