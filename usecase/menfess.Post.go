package usecase

import "github.com/aeramu/qiup-server/entity"

func (i *interactor) MenfessPost(id string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(id)
	return post
}

func (i *interactor) MenfessPostFeed(first int, after string) []entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID("", first, after, false)
	return postList
}

func (i *interactor) MenfessPostChild(parentID string, first int, after string) []entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByParentID(parentID, first, after, true)
	return postList
}

func (i *interactor) MenfessPostRooms(roomIDs []string, first int, after string) []entity.MenfessPost {
	postList := i.menfessPostRepo.GetDataListByRoomIDs(roomIDs, first, after, false)
	return postList
}

func (i *interactor) PostMenfessPost(name string, avatar string, body string, parentID string, repostID string, roomID string) entity.MenfessPost {
	post := i.menfessPostRepo.PutData(name, avatar, body, parentID, repostID, roomID)
	return post
}

func (i *interactor) UpvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	if post.IsDownvoted(accountID) {
		exist := post.Downvote(accountID)
		i.menfessPostRepo.UpdateDownvoterIDs(postID, accountID, exist)
	}
	exist := post.Upvote(accountID)
	i.menfessPostRepo.UpdateUpvoterIDs(postID, accountID, exist)
	return post
}

func (i *interactor) DownvoteMenfessPost(accountID string, postID string) entity.MenfessPost {
	post := i.menfessPostRepo.GetDataByID(postID)
	if post.IsUpvoted(accountID) {
		exist := post.Upvote(accountID)
		i.menfessPostRepo.UpdateUpvoterIDs(postID, accountID, exist)
	}
	exist := post.Downvote(accountID)
	i.menfessPostRepo.UpdateDownvoterIDs(postID, accountID, exist)
	return post
}
