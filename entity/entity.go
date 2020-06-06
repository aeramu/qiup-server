package entity

type Account struct{
	ID string `bson:"_id"`
	Email string
	Password string
}

type ShareAccount struct{
	ID string `bson:"_id"`
	Username string
	ShareProfile ShareProfile `bson:"shareProfile"`
}

type ShareProfile struct{
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
}