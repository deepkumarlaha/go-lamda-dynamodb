// pkg/dynamodb.go
package pkg

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamoDBClient *dynamodb.DynamoDB
var tableName = "UserTable"

func init() {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		log.Printf("Error creating AWS session: %v", err)
		return
	}
	dynamoDBClient = dynamodb.New(awsSession)
}

// CreateUser adds a new user to DynamoDB
func CreateUser(user User) error {
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

// GetUserByID retrieves a user from DynamoDB based on userID
func GetUserByID(userID string) (User, error) {
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

// UpdateUser updates user information in DynamoDB
func UpdateUser(user User) error {
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

// DeleteUser removes a user from DynamoDB based on userID
func DeleteUser(userID string) error {
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
