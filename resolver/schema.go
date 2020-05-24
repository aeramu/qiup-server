package resolver

type Resolver struct{}

var Schema = `
  	schema{
		query: Query
  	}
	
  	type Query{
		hello: String!
	}
`

func (r *Resolver) Hello()(string){
	return "Hello world!"
}