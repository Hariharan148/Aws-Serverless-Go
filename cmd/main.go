package main

import (
	"os"
	""
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/service/lambda"
	"github.com/aws/aws-sdk-go/aws/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-lambda-go/events"
	"github.com/Hariharan148/Aws-Serverless-Go/pkg/handlers"
)

var (
	dynalient dynamodbiface.DynamoDBAPI
)


func main(){
	const region = os.Getenv("AWS-REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}

	dynaClient = dynamodb.New(awsSession)

	lambda.Start(handler(dynaClient))
}

const tableName = "LambdaInGo"

func handler(req events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse, error){
	switch req.HTTPMethod{
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledReq()
	}
}