package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		account(id: ID!): Account
		me: Account!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, username: String!, password: String!): String!
		editProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): Account!
		uploadImage(directory: String!): String!
		isEmailAvailable(email: String!): Boolean!
		isUsernameAvailable(username: String!): Boolean!
	}
	type Account{
		id: ID!
		email: String!
		username: String!
		profile: Profile!
	}
	type Profile{
		name: String!
		bio: String!
		profilePhoto: String!
		coverPhoto: String!
	}
`