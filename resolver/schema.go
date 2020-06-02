package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		isEmailAvailable(email: String!): Boolean!
		account(id: ID!): Account
		me: Account!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!
		editProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): Account!
		uploadImage(directory: String!): String!
		isUsernameAvailable(username: String!): Boolean!
	}
	type Account{
		id: ID!
		email: String!
	}
	type Profile{
		id: ID!
		username: String!
		name: String!
		bio: String!
		profilePhoto: String!
		coverPhoto: String!
	}
`