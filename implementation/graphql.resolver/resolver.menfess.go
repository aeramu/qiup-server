package resolver

import (
	"github.com/aeramu/qiup-server/entity"
	"github.com/graph-gophers/graphql-go"
)

// MenfessPostResolver graphql
type MenfessPostResolver struct {
	post entity.MenfessPost
	pr   *Resolver
}

// ID graphql
func (r *MenfessPostResolver) ID() graphql.ID {
	return graphql.ID(r.post.ID())
}

// Timestamp graphql
func (r *MenfessPostResolver) Timestamp() int32 {
	return int32(r.post.Timestamp())
}

// Name graphql
func (r *MenfessPostResolver) Name() string {
	return r.post.Name()
}

// Avatar graphql
func (r *MenfessPostResolver) Avatar() string {
	return r.post.Avatar()
}

// Body graphql
func (r *MenfessPostResolver) Body() string {
	return r.post.Body()
}

//Room graphql
func (r *MenfessPostResolver) Room() string {
	room := r.pr.Interactor.MenfessRoom(r.post.RoomID())
	if room == nil {
		return "General"
	}
	return room.Name()
}

// ReplyCount graphql
func (r *MenfessPostResolver) ReplyCount() int32 {
	return int32(r.post.ReplyCount())
}

// UpvoteCount graphql
func (r *MenfessPostResolver) UpvoteCount() int32 {
	return int32(r.post.UpvoteCount())
}

// DownvoteCount graphql
func (r *MenfessPostResolver) DownvoteCount() int32 {
	return int32(r.post.DownvoteCount())
}

//Upvoted bool
func (r *MenfessPostResolver) Upvoted() bool {
	accountID := r.pr.Context.Value("request").(map[string]string)["id"]
	return r.post.IsUpvoted(accountID)
}

//Downvoted bool
func (r *MenfessPostResolver) Downvoted() bool {
	accountID := r.pr.Context.Value("request").(map[string]string)["id"]
	return r.post.IsDownvoted(accountID)
}

// Parent graphql
func (r *MenfessPostResolver) Parent() *MenfessPostResolver {
	post := r.pr.Interactor.MenfessPost(r.post.ParentID())
	if post == nil {
		return nil
	}
	return &MenfessPostResolver{post, r.pr}
}

//Repost graphql
func (r *MenfessPostResolver) Repost() *MenfessPostResolver {
	post := r.pr.Interactor.MenfessPost(r.post.RepostID())
	if post == nil {
		return nil
	}
	return &MenfessPostResolver{post, r.pr}
}

// Child graphql
func (r *MenfessPostResolver) Child(args struct {
	First  *int32
	After  *graphql.ID
	Before *graphql.ID
	Sort   *int32
}) *MenfessPostConnectionResolver {
	first := 20
	if args.First != nil {
		first = int(*args.First)
	}
	after := "000000000000000000000000"
	postList := r.pr.Interactor.MenfessPostChild(r.post.ID(), first, after)
	return &MenfessPostConnectionResolver{postList, r.pr}
}

// MenfessPostConnectionResolver graphql
type MenfessPostConnectionResolver struct {
	menfessPostList []entity.MenfessPost
	pr              *Resolver
}

// Edges graphql
func (r *MenfessPostConnectionResolver) Edges() []*MenfessPostResolver {
	var menfessPostResolverList []*MenfessPostResolver
	for _, post := range r.menfessPostList {
		menfessPostResolverList = append(menfessPostResolverList, &MenfessPostResolver{post, r.pr})
	}
	return menfessPostResolverList
}

// PageInfo graphql
func (r *MenfessPostConnectionResolver) PageInfo() *PageInfoResolver {
	var nodeList []node
	for _, node := range r.menfessPostList {
		nodeList = append(nodeList, node)
	}
	return &PageInfoResolver{nodeList}
}

// MenfessPost graphql
func (r *Resolver) MenfessPost(args struct {
	ID graphql.ID
}) *MenfessPostResolver {
	post := r.Interactor.MenfessPost(string(args.ID))
	return &MenfessPostResolver{post, r}
}

// MenfessPostList graphql
func (r *Resolver) MenfessPostList(args struct {
	First *int32
	After *graphql.ID
	Sort  *bool
}) *MenfessPostConnectionResolver {
	first := 20
	if args.First != nil {
		first = int(*args.First)
	}
	after := "ffffffffffffffffffffffff"
	if args.After != nil {
		after = string(*args.After)
	}
	postList := r.Interactor.MenfessPostFeed(first, after)
	return &MenfessPostConnectionResolver{postList, r}
}

//MenfessPostRooms graphql
func (r *Resolver) MenfessPostRooms(args struct {
	IDs   []graphql.ID
	First *int32
	After *graphql.ID
}) *MenfessPostConnectionResolver {
	first := 20
	if args.First != nil {
		first = int(*args.First)
	}
	after := "ffffffffffffffffffffffff"
	if args.After != nil {
		after = string(*args.After)
	}
	var roomIDs []string
	for _, id := range args.IDs {
		roomIDs = append(roomIDs, string(id))
	}
	postList := r.Interactor.MenfessPostRooms(roomIDs, first, after)
	return &MenfessPostConnectionResolver{postList, r}
}

// PostMenfessPost graphql
func (r *Resolver) PostMenfessPost(args struct {
	Name     string
	Avatar   string
	Body     string
	ParentID *graphql.ID
	RepostID *graphql.ID
	RoomID   *graphql.ID
}) *MenfessPostResolver {
	parentID := ""
	if args.ParentID != nil {
		parentID = string(*args.ParentID)
	}
	repostID := ""
	if args.RepostID != nil {
		repostID = string(*args.RepostID)
	}
	roomID := ""
	if args.RoomID != nil {
		roomID = string(*args.RoomID)
	}
	post := r.Interactor.PostMenfessPost(args.Name, args.Avatar, args.Body, parentID, repostID, roomID)
	return &MenfessPostResolver{post, r}
}

//UpvoteMenfessPost graphql
func (r *Resolver) UpvoteMenfessPost(args struct {
	PostID graphql.ID
}) *MenfessPostResolver {
	accountID := r.Context.Value("request").(map[string]string)["id"]
	post := r.Interactor.UpvoteMenfessPost(accountID, string(args.PostID))
	return &MenfessPostResolver{post, r}
}

//DownvoteMenfessPost graphql
func (r *Resolver) DownvoteMenfessPost(args struct {
	PostID graphql.ID
}) *MenfessPostResolver {
	accountID := r.Context.Value("request").(map[string]string)["id"]
	post := r.Interactor.DownvoteMenfessPost(accountID, string(args.PostID))
	return &MenfessPostResolver{post, r}
}

// MenfessAvatarList graphql
func (r *Resolver) MenfessAvatarList() []string {
	avatarList := []string{
		"https://qiup-image.s3.amazonaws.com/avatar/avatar.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/batman.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/spiderman.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/saitama.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/kaonashi.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/mrbean.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/upin.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/ipin.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/einstein.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/monalisa.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/ronald.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/1cokelat.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/2merah.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/3vermilion.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/4oranye.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/5oranye_muda.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/6kuning.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/7hijau.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/8hijau_daun.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/9toska.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/10biru.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/11biru_tua.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/12blue-violet.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/13ungu.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/14red-violet.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/15magenta.jpg",
		"https://qiup-image.s3.amazonaws.com/avatar/16pink.jpg",
	}
	return avatarList
}

//MenfessRoomList graphql
func (r *Resolver) MenfessRoomList() *MenfessRoomConnectionResolver {
	// room1 := entity.MenfessPostConstructor{
	// 	ID:   "5ef89baaec8ff2af8b9934c1",
	// 	Name: "ITB",
	// }.New()
	// room2 := entity.MenfessPostConstructor{
	// 	ID:   "5efadcbdc0245024fd758d02",
	// 	Name: "UI",
	// }.New()
	// roomList := []entity.MenfessRoom{room1, room2}
	roomList := r.Interactor.MenfessRoomList()
	return &MenfessRoomConnectionResolver{roomList, r}
}

//MenfessRoomResolver graphql
type MenfessRoomResolver struct {
	room entity.MenfessRoom
	pr   *Resolver
}

//ID get
func (r *MenfessRoomResolver) ID() graphql.ID {
	return graphql.ID(r.room.ID())
}

//Name get
func (r *MenfessRoomResolver) Name() string {
	return r.room.Name()
}

// MenfessRoomConnectionResolver graphql
type MenfessRoomConnectionResolver struct {
	menfessRoomList []entity.MenfessRoom
	pr              *Resolver
}

// Edges graphql
func (r *MenfessRoomConnectionResolver) Edges() []*MenfessRoomResolver {
	var menfessRoomResolverList []*MenfessRoomResolver
	for _, room := range r.menfessRoomList {
		menfessRoomResolverList = append(menfessRoomResolverList, &MenfessRoomResolver{room, r.pr})
	}
	return menfessRoomResolverList
}

// PageInfo graphql
func (r *MenfessRoomConnectionResolver) PageInfo() *PageInfoResolver {
	var nodeList []node
	for _, node := range r.menfessRoomList {
		nodeList = append(nodeList, node)
	}
	return &PageInfoResolver{nodeList}
}
