package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		myAccount: Account!
		isEmailAvailable(email: String!): Boolean!

		profile(id: ID!): Profile!
		myProfile: Profile!
		isUsernameAvailable(username: String!): Boolean!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!

		editProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): Profile!
		uploadImage(directory: String!): String!
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