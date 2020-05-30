package main

import(
  	"context"
  	"encoding/json"
  	"github.com/aeramu/qiup-server/resolver"
  	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
 	"github.com/graph-gophers/graphql-go"
)

func main(){
  	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error){
  	//convert request body to json
  	var parameter struct{
    	Query string
      	OperationName string
      	Variables map[string]interface{}
  	}
  	json.Unmarshal([]byte(request.Body), &parameter)

  	//add token from header
  	ctxWithToken := context.WithValue(ctx, "token", request.Headers["token"])

  	//graphql execution
  	schema := graphql.MustParseSchema(resolver.Schema, &resolver.Resolver{})
  	response := schema.Exec(ctxWithToken, parameter.Query, parameter.OperationName, parameter.Variables)
  	responseJSON,_ := json.Marshal(response)

  	//response
  	return events.APIGatewayProxyResponse{
		Headers: map[string]string {
			"Access-Control-Allow-Origin" : "*",
		},
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}
