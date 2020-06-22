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

		justPost(id: ID!): JustPost!
		justPostList(first: Int=20, after: ID="ffffffffffffffffffffffff"): JustPostConnection!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!
		isEmailAvailable(email: String!): Boolean!

		setShareUsername(username: String!): String!
		setShareProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): ShareAccount!
		isUsernameAvailable(username: String!): Boolean!

		postSharePost(body: String!): SharePost!

		postJustPost(name: String!, avatar: String!, body: String!, parentID: ID=""): JustPost!

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
	type JustPost{
		id: ID!
		timestamp: Int!
		parent: JustPost
		name: String!
		avatar: String!
		body: String!
		child(first: Int=20, after: ID="ffffffffffffffffffffffff"): JustPostConnection!
		replyCount: Int!
	}
	type JustPostConnection{
		edges: [JustPost]!
		pageInfo: PageInfo!
	}
	type PageInfo{
		startCursor: ID
		endCursor: ID
	}
`