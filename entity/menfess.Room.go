package entity

//MenfessRoom interface
type MenfessRoom interface {
	ID() string
	Name() string
}

//MenfessRoomConstructor constructor
type MenfessRoomConstructor struct {
	ID   string
	Name string
}

//New constructor
func (c MenfessRoomConstructor) New() MenfessRoom {
	return &menfessRoom{
		id:   c.ID,
		name: c.Name,
	}
}

type menfessRoom struct {
	id   string
	name string
}

func (mr *menfessRoom) ID() string {
	return mr.id
}

func (mr *menfessRoom) Name() string {
	return mr.name
}
