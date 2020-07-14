package usecase

import "github.com/aeramu/qiup-server/entity"

func (i *interactor) MenfessPost(id string) entity.MenfessPost {
	post := i.menfessRepo.GetPostByID(id)
	return post
}

func (i *interactor) MenfessPostFeed(first int, after string) []entity.MenfessPost {
	postList := i.menfessRepo.GetPostListByParentID("", first, after, false)
	return postList
}

func (i *interactor) MenfessPostChild(parentID string, first int, after string) []entity.MenfessPost {
	postList := i.menfessRepo.GetPostListByParentID(parentID, first, after, true)
	return postList
}

func (i *interactor) MenfessPostRooms(roomIDs []string, first int, after string) []entity.MenfessPost {
	postList := i.menfessRepo.GetPostListByRoomIDs(roomIDs, first, after, false)
	return postList
}

func (i *interactor) PostMenfessPost(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost {
	post := i.menfessRepo.PutPost(name, avatar, body, parentID, repostID, roomID)
	return post
}

func (i *interactor) UpvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessRepo.GetPostByID(postID)
	if post.IsDownvoted(accountID) {
		exist := post.Downvote(accountID)
		i.menfessRepo.UpdateDownvoterIDs(postID, accountID, exist)
	}
	exist := post.Upvote(accountID)
	i.menfessRepo.UpdateUpvoterIDs(postID, accountID, exist)
	return post
}

func (i *interactor) DownvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessRepo.GetPostByID(postID)
	if post.IsUpvoted(accountID) {
		exist := post.Upvote(accountID)
		i.menfessRepo.UpdateUpvoterIDs(postID, accountID, exist)
	}
	exist := post.Downvote(accountID)
	i.menfessRepo.UpdateDownvoterIDs(postID, accountID, exist)
	return post
}
