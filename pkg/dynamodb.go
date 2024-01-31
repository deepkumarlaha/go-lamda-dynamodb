// pkg/dynamodb.go
package pkg

import (
	"log"
	"strconv"

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
			"name":    {S: aws.String(user.Name)},
			"email":   {S: aws.String(user.Email)},
			"phone":   {N: aws.String(strconv.Itoa(user.Phone))},
			"gender":  {S: aws.String(user.Gender)},
			"address": {S: aws.String(user.Address)},
			"state":   {S: aws.String(user.State)},
			"country": {S: aws.String(user.Country)},
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

// GetUserByemail retrieves a user from DynamoDB based on email
func GetUserByemail(email string) (User, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {S: aws.String(email)},
		},
	}

	result, err := dynamoDBClient.GetItem(params)
	if err != nil {
		log.Printf("Error getting item from DynamoDB: %v", err)
		return User{}, err
	}

	user := User{
		Name:    *result.Item["name"].S,
		Email:   *result.Item["email"].S,
		Gender:  *result.Item["gender"].S,
		Address: *result.Item["address"].S,
		State:   *result.Item["state"].S,
		Country: *result.Item["country"].S,
		// Add other attributes as needed...
	}

	// Check if "phone" attribute exists before attempting conversion
	if phoneAttr, ok := result.Item["phone"]; ok && phoneAttr.N != nil {
		phone, err := strconv.Atoi(*phoneAttr.N)
		if err != nil {
			log.Printf("Error converting phone to int: %v", err)
			return User{}, err
		}
		user.Phone = phone
	}

	return user, nil
}

// UpdateUser updates user information in DynamoDB
func UpdateUser(user User) error {
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {S: aws.String(user.Email)},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {S: aws.String(user.Name)},
			// Add other attributes as needed...
		},
		UpdateExpression: aws.String("SET name = :n"),
	}

	_, err := dynamoDBClient.UpdateItem(params)
	if err != nil {
		log.Printf("Error updating item in DynamoDB: %v", err)
		return err
	}

	return nil
}

// DeleteUser removes a user from DynamoDB based on email
func DeleteUser(email string) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {S: aws.String(email)},
		},
	}

	_, err := dynamoDBClient.DeleteItem(params)
	if err != nil {
		log.Printf("Error deleting item from DynamoDB: %v", err)
		return err
	}

	return nil
}
