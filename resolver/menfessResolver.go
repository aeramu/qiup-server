package resolver

import (
	"github.com/aeramu/qiup-server/entity"
	"github.com/aeramu/qiup-server/repository"
	"github.com/graph-gophers/graphql-go"
)

// MenfessPostResolver graphql
type MenfessPostResolver struct {
	post *entity.MenfessPost
}

// ID graphql
func (r *MenfessPostResolver) ID() graphql.ID {
	return graphql.ID(r.post.ID.Hex())
}

// Timestamp graphql
func (r *MenfessPostResolver) Timestamp() int32 {
	return int32(r.post.Timestamp())
}

// Name graphql
func (r *MenfessPostResolver) Name() string {
	return r.post.Name
}

// Avatar graphql
func (r *MenfessPostResolver) Avatar() string {
	return r.post.Avatar
}

// Body graphql
func (r *MenfessPostResolver) Body() string {
	return r.post.Body
}

// ReplyCount graphql
func (r *MenfessPostResolver) ReplyCount() int32 {
	return r.post.ReplyCount
}

// UpvoteCount graphql
func (r *MenfessPostResolver) UpvoteCount() int32 {
	return r.post.UpvoteCount
}

// DownvoteCount graphql
func (r *MenfessPostResolver) DownvoteCount() int32 {
	return r.post.DownvoteCount
}

// Parent graphql
func (r *MenfessPostResolver) Parent() *MenfessPostResolver {
	menfessPostRepository := repository.NewMenfessPostRepository()
	post := menfessPostRepository.GetDataByIndex("_id", r.post.ParentID)
	if post == nil {
		return nil
	}
	return &MenfessPostResolver{post}
}

// Child graphql
func (r *MenfessPostResolver) Child(args struct {
	First  *int32
	After  *graphql.ID
	Before *graphql.ID
	Sort   *int32
}) *MenfessPostConnectionResolver {
	first := int32(20)
	if args.First != nil {
		first = *args.First
	}
	sort := int32(1)
	if args.Sort != nil {
		sort = *args.Sort
	}
	after := entity.ID("")
	if sort == -1 {
		after = entity.ID("ffffffffffffffffffffffff")
	}
	if args.After != nil {
		after = entity.ID(string(*args.After))
	}
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostList := menfessPostRepository.GetDataListByIndex("parentID", r.post.ID, first, after, sort)
	return &MenfessPostConnectionResolver{menfessPostList}
}

// MenfessPostConnectionResolver graphql
type MenfessPostConnectionResolver struct {
	menfessPostList []*entity.MenfessPost
}

// Edges graphql
func (r *MenfessPostConnectionResolver) Edges() []*MenfessPostResolver {
	var menfessPostResolverList []*MenfessPostResolver
	for _, post := range r.menfessPostList {
		menfessPostResolverList = append(menfessPostResolverList, &MenfessPostResolver{post})
	}
	return menfessPostResolverList
}

// PageInfo graphql
func (r *MenfessPostConnectionResolver) PageInfo() *PageInfoResolver {
	var nodeList []entity.Node
	for _, node := range r.menfessPostList {
		nodeList = append(nodeList, node)
	}
	return &PageInfoResolver{nodeList}
}

// MenfessPost graphql
func (r *Resolver) MenfessPost(args struct {
	ID graphql.ID
}) *MenfessPostResolver {
	menfessPostRepository := repository.NewMenfessPostRepository()
	post := menfessPostRepository.GetDataByIndex("_id", entity.ID(string(args.ID)))
	return &MenfessPostResolver{post}
}

// MenfessPostList graphql
func (r *Resolver) MenfessPostList(args struct {
	First  *int32
	After  *graphql.ID
	Before *graphql.ID
	Sort   *int32
}) *MenfessPostConnectionResolver {
	first := int32(20)
	if args.First != nil {
		first = *args.First
	}
	sort := int32(-1)
	if args.Sort != nil {
		sort = *args.Sort
	}
	after := entity.ID("")
	if sort == -1 {
		after = entity.ID("ffffffffffffffffffffffff")
	}
	if args.After != nil {
		after = entity.ID(string(*args.After))
	}
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostList := menfessPostRepository.GetDataListByIndex("parentID", entity.ID(""), first, after, sort)
	return &MenfessPostConnectionResolver{menfessPostList}
}

// PostMenfessPost graphql
func (r *Resolver) PostMenfessPost(args struct {
	Name     string
	Avatar   string
	Body     string
	ParentID *graphql.ID
}) *MenfessPostResolver {
	post := entity.NewMenfessPost(entity.ID("")).
		SetName(args.Name).
		SetAvatar(args.Avatar).
		SetBody(args.Body).
		SetParentID(entity.ID(string(*args.ParentID)))
	menfessPostRepository := repository.NewMenfessPostRepository()
	menfessPostRepository.PutData(post)
	return &MenfessPostResolver{post}
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
