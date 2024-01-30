// pkg/handler.go
package pkg

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// User represents a user entity
type User struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	// Add other attributes as needed...
}

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid JSON input"}, nil
	}

	err = CreateUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// GetUserHandler retrieves user information based on userID
func GetUserHandler(userID string) (events.APIGatewayProxyResponse, error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error getting user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// GetAllUsersHandler retrieves information for all users
func GetAllUsersHandler() (events.APIGatewayProxyResponse, error) {
	// Implement logic to retrieve all users
	// ...
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "List of users"}, nil
}

// UpdateUserHandler updates user information based on the request
func UpdateUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid JSON input"}, nil
	}

	err = UpdateUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error updating user"}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// DeleteUserHandler removes a user based on userID
func DeleteUserHandler(userID string) (events.APIGatewayProxyResponse, error) {
	err := DeleteUser(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting user"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "User deleted successfully"}, nil
}
