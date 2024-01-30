// main.go
package main

import (
	"context"

	"lambda/project/pkg"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return pkg.CreateUserHandler(request)
	case "GET":
		if userID, ok := request.PathParameters["id"]; ok {
			return pkg.GetUserHandler(userID)
		} else {
			return pkg.GetAllUsersHandler()
		}
	case "PUT":
		return pkg.UpdateUserHandler(request)
	case "DELETE":
		if userID, ok := request.PathParameters["id"]; ok {
			return pkg.DeleteUserHandler(userID)
		} else {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request"}, nil
		}
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request"}, nil
	}
}

func main() {
	lambda.Start(handler)
}
