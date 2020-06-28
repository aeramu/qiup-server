package entity

//MenfessPost interface
type MenfessPost interface {
	ID() string
	Timestamp() int
	Name() string
	Avatar() string
	Body() string
	UpvoterIDs() map[string]bool
	DownvoterIDs() map[string]bool
	UpvoteCount() int
	DownvoteCount() int
	ReplyCount() int
	ParentID() string
	IsUpvoted(accountID string) bool
	IsDownvoted(accountID string) bool
	Upvote(accountID string) bool
	Downvote(accountID string) bool
}

//MenfessPostConstructor struct
type MenfessPostConstructor struct {
	ID           string
	Timestamp    int
	Name         string
	Avatar       string
	Body         string
	UpvoterIDs   map[string]bool
	DownvoterIDs map[string]bool
	ReplyCount   int
	ParentID     string
}

//New construtor
func (c MenfessPostConstructor) New() MenfessPost {
	if c.UpvoterIDs == nil {
		c.UpvoterIDs = map[string]bool{}
	}
	if c.DownvoterIDs == nil {
		c.DownvoterIDs = map[string]bool{}
	}
	if c.ReplyCount == 0 {
		c.ReplyCount = 0
	}
	return &menfessPost{
		id:           c.ID,
		timestamp:    c.Timestamp,
		name:         c.Name,
		avatar:       c.Avatar,
		body:         c.Body,
		upvoterIDs:   c.UpvoterIDs,
		downvoterIDs: c.DownvoterIDs,
		replyCount:   c.ReplyCount,
		parentID:     c.ParentID,
	}
}

type menfessPost struct {
	id           string
	timestamp    int
	name         string
	avatar       string
	body         string
	upvoterIDs   map[string]bool
	downvoterIDs map[string]bool
	replyCount   int
	parentID     string
}

func (mp *menfessPost) Upvote(accountID string) bool {
	voted := mp.IsUpvoted(accountID)
	if !voted {
		mp.upvoterIDs[accountID] = true
	} else {
		delete(mp.upvoterIDs, accountID)
	}
	return voted
}
func (mp *menfessPost) Downvote(accountID string) bool {
	voted := mp.IsDownvoted(accountID)
	if !voted {
		mp.downvoterIDs[accountID] = true
	} else {
		delete(mp.downvoterIDs, accountID)
	}
	return voted
}
func (mp *menfessPost) IsUpvoted(accountID string) bool {
	_, ok := mp.upvoterIDs[accountID]
	return ok
}
func (mp *menfessPost) IsDownvoted(accountID string) bool {
	_, ok := mp.downvoterIDs[accountID]
	return ok
}
func (mp *menfessPost) ParentID() string {
	return mp.parentID
}
func (mp *menfessPost) ReplyCount() int {
	return mp.replyCount
}
func (mp *menfessPost) DownvoteCount() int {
	return len(mp.downvoterIDs)
}
func (mp *menfessPost) UpvoteCount() int {
	return len(mp.upvoterIDs)
}
func (mp *menfessPost) UpvoterIDs() map[string]bool {
	return mp.upvoterIDs
}
func (mp *menfessPost) DownvoterIDs() map[string]bool {
	return mp.downvoterIDs
}
func (mp *menfessPost) Body() string {
	return mp.body
}
func (mp *menfessPost) Avatar() string {
	return mp.avatar
}
func (mp *menfessPost) Name() string {
	return mp.name
}
func (mp *menfessPost) Timestamp() int {
	return mp.timestamp
}
func (mp *menfessPost) ID() string {
	return mp.id
}
