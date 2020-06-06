package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		myAccount: Account!

		profile(id: ID!): Profile!
		myProfile: Profile!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!
		isEmailAvailable(email: String!): Boolean!

		setProfile(username: String!, name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): Profile!
		isUsernameAvailable(username: String!): Boolean!
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