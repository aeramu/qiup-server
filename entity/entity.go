package entity

type Account struct{
	ID string `bson:"_id"`
	Email string
	Username string
	Password string
	Profile *Profile
}

type Profile struct{
	Name string
	Bio string
}