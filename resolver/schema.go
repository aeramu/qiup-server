package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
		mutation: Mutation
  	}
  	type Query{
		myAccount: Account!

		shareAccount(id: ID!): ShareAccount!
		myShareAccount: ShareAccount!

		sharePost(id: ID!): SharePost!
		sharePostList: [SharePost]!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!
		isEmailAvailable(email: String!): Boolean!

		setShareUsername(username: String!): String!
		setShareProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): ShareAccount!
		isUsernameAvailable(username: String!): Boolean!

		postSharePost(body: String!): SharePost!

		uploadImage(directory: String!): String!
	}
	type Account{
		id: ID!
		email: String!
	}
	type ShareAccount{
		id: ID!
		username: String!
		name: String!
		bio: String!
		profilePhoto: String!
		coverPhoto: String!
	}
	type SharePost{
		id: ID!
		timestamp: Int!
		account: ShareAccount!
		body: String!
	}
`