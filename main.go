// main.go
package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	dynamoDBClient *dynamodb.DynamoDB
	tableName      string = "UserTable"
)

// User represents a user entity
type User struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	// Add other attributes as needed...
}

func init() {
	// Initialize DynamoDB client
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		log.Printf("Error creating AWS session: %v", err)
		return
	}
	dynamoDBClient = dynamodb.New(awsSession)
}

func createUser(user User) error {
	params := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(user.UserID)},
			"UserName": {S: aws.String(user.UserName)},
			// Add other attributes as needed...
		},
	}

	_, err := dynamoDBClient.PutItem(params)
	if err != nil {
		log.Printf("Error putting item into DynamoDB: %v", err)
		return err
	}

	return nil
}

func getUserByID(userID string) (User, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(userID)},
		},
	}

	result, err := dynamoDBClient.GetItem(params)
	if err != nil {
		log.Printf("Error getting item from DynamoDB: %v", err)
		return User{}, err
	}

	user := User{
		UserID:   *result.Item["UserID"].S,
		UserName: *result.Item["UserName"].S,
		// Retrieve other attributes as needed...
	}

	return user, nil
}

func updateUser(user User) error {
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(user.UserID)},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {S: aws.String(user.UserName)},
			// Add other attributes as needed...
		},
		UpdateExpression: aws.String("SET UserName = :n"),
	}

	_, err := dynamoDBClient.UpdateItem(params)
	if err != nil {
		log.Printf("Error updating item in DynamoDB: %v", err)
		return err
	}

	return nil
}

func deleteUser(userID string) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(userID)},
		},
	}

	_, err := dynamoDBClient.DeleteItem(params)
	if err != nil {
		log.Printf("Error deleting item from DynamoDB: %v", err)
		return err
	}

	return nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return createUserHandler(request)
	case "GET":
		if userID, ok := request.PathParameters["id"]; ok {
			return getUserHandler(userID)
		} else {
			return getAllUsersHandler()
		}
	case "PUT":
		return updateUserHandler(request)
	case "DELETE":
		if userID, ok := request.PathParameters["id"]; ok {
			return deleteUserHandler(userID)
		} else {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request"}, nil
		}
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request"}, nil
	}
}

func createUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid JSON input"}, nil
	}

	err = createUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

func getUserHandler(userID string) (events.APIGatewayProxyResponse, error) {
	user, err := getUserByID(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error getting user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

func getAllUsersHandler() (events.APIGatewayProxyResponse, error) {
	// Implement logic to retrieve all users
	// ...
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "List of users"}, nil
}

func updateUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid JSON input"}, nil
	}

	err = updateUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error updating user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

func deleteUserHandler(userID string) (events.APIGatewayProxyResponse, error) {
	err := deleteUser(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting user"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "User deleted successfully"}, nil
}

func main() {
	lambda.Start(handler)
}
