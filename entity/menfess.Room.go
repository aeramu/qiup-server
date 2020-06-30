package entity

//MenfessRoom interface
type MenfessRoom interface {
	ID() string
	Name() string
	Posts() string
}

type menfessRoom struct {
	id    string
	name  string
	posts []MenfessPost
}

func (mr *menfessRoom) ID() string {
	return mr.id
}

func (mr *menfessRoom) Name() string {
	return mr.name
}

func (mr *menfessRoom) Posts() []MenfessPost {
	return mr.posts
}
