package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
	
  	type Query{
		hello: String!
		login(email: String = "", username: String = "", password: String!): String!
		me: Account!
		account(id: ID!): Account!
	}

	type Mutation{
		register(email: String!, username: String!, password: String!): String!
		changeProfile(name: String!, bio: String!): Profile!
	}

	type Account{
		accountID: ID!
		email: String!
		username: String!
		password: String!
		profile: Profile!
	}

	type Profile{
		name: String
		bio: String
	}
`

func (r *Resolver) Hello()(string){
	return "Hello world!"
}