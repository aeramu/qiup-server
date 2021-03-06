package resolver

import (
	"context"

	"github.com/aeramu/qiup-server/usecase"
)

//Resolver graphql
type Resolver struct {
	Interactor usecase.Interactor
	Context    context.Context
}

//Schema grahql
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

		menfessPost(id: ID!): MenfessPost!
		menfessPostList(first: Int, after: ID, sort: Int): MenfessPostConnection!
		menfessPostRooms(ids: [ID!]!, first: Int, after: ID): MenfessPostConnection!
		menfessRoomList: MenfessRoomConnection!
		menfessAvatarList: [String!]!
	}
	type Mutation{
		login(email: String!, password: String!): String!
		register(email: String!, password: String!): String!
		isEmailAvailable(email: String!): Boolean!

		setShareUsername(username: String!): String!
		setShareProfile(name: String!, bio: String!, profilePhoto: String!, coverPhoto: String!): ShareAccount!
		isUsernameAvailable(username: String!): Boolean!

		postSharePost(body: String!): SharePost!

		postMenfessPost(name: String!, avatar: String!, body: String!, parentID: ID, repostID: ID, roomID: ID): MenfessPost!
		upvoteMenfessPost(postID: ID!): MenfessPost!
		downvoteMenfessPost(postID: ID!): MenfessPost!

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
	type MenfessPost{
		id: ID!
		timestamp: Int!
		name: String!
		avatar: String!
		body: String!
		replyCount: Int!
		upvoteCount: Int!
		downvoteCount: Int!
		upvoted: Boolean!
		downvoted: Boolean!
		parent: MenfessPost
		repost: MenfessPost
		child(first: Int, after: ID, before: ID, sort: Int): MenfessPostConnection!
		room: String!
	}
	type MenfessRoom{
		id: ID!
		name: String!
	}
	type MenfessPostConnection{
		edges: [MenfessPost]!
		pageInfo: PageInfo!
	}
	type MenfessRoomConnection{
		edges: [MenfessRoom]!
		pageInfo: PageInfo!
	}
	type PageInfo{
		startCursor: ID
		endCursor: ID
	}
`
