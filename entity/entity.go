package entity

type Account struct{
	ID string `bson:"_id"`
	Email string
	Password string
}

type Profile struct{
	ID string `bson:"_id"`
	Username string
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
}