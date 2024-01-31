// pkg/handler.go
package pkg

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// User represents a user entity
// User represents a user entity
type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	State   string `json:"state"`
	Country string `json:"country"`
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
func GetUserHandler(email string) (events.APIGatewayProxyResponse, error) {
	user, err := GetUserByemail(email)
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

// DeleteUserHandler removes a user based on email
func DeleteUserHandler(email string) (events.APIGatewayProxyResponse, error) {
	err := DeleteUser(email)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting user"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "User deleted successfully"}, nil
}
