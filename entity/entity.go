package entity

type Account struct{
	ID string `bson:"_id"`
	Email string
	Password string
}

type Profile struct{
	Name string
	Bio string
	ProfilePhoto string
	CoverPhoto string
}