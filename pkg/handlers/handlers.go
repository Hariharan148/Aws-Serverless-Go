package handlers

import (
	"github.com/aws/aws-sdk-go/aws/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-lambda-go/events"

)

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI )(events.APIGatewayProxyResponse, error){

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI )(events.APIGatewayProxyResponse, error){

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI )(events.APIGatewayProxyResponse, error){

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI )(events.APIGatewayProxyResponse, error){

}

func UnhandledReq()(){

}