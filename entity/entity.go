package entity

type Account struct{
	ID string
	Email string
	Username string
	Password string
	Profile Profile
}

type Profile struct{
	Name string
	Bio string
}