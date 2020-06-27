package entity

//MenfessPost entity
type MenfessPost struct {
	ID            string
	Timestamp     int
	Name          string
	Avatar        string
	Body          string
	UpvoterIDs    []string
	DownvoterIDs  []string
	UpvoteCount   int
	DownvoteCount int
	ReplyCount    int
	ParentID      string
}

//GetID interface node
func (m *MenfessPost) GetID() string {
	return m.ID
}
