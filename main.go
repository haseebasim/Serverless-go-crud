package main

import (
	"encoding/json"
	"fmt"
	"go-lambda-postgres/middleware"
	"go-lambda-postgres/models"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BodyPostRequest struct {
	RequestFirstName string `json:"firstname"`
	RequestLastName  string `json:"lastname"`
}

type BodyGetRequest struct {
	RequestId string `json:"id"`
}

type PostResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(request.HTTPMethod)
	fmt.Println(request.Body)
	fmt.Println(request.Path)

	switch request.HTTPMethod {
	case "GET":
		if request.Path == "/" {
			users, dbError := middleware.GetAllUsers()
			if dbError != nil {
				log.Fatalf("Unable to convert the string into int.  %v", dbError)
			}
			response, err := json.Marshal(users)
			if err != nil {
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
			}
			return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

		} else {

			id, intParseError := strconv.Atoi(request.PathParameters["id"])

			if intParseError != nil {
				log.Fatalf("Unable to convert the string into int.  %v", intParseError)
			}

			user, dbError := middleware.GetUser(int64(id))
			if dbError != nil {
				log.Fatalf("Unable to convert the string into int.  %v", dbError)
			}
			response, err := json.Marshal(user)
			if err != nil {
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
			}
			return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
		}
	case "POST":
		fmt.Println("post running")

		now := time.Now()

		userModel := models.User{
			FirstName:    request.QueryStringParameters["firstname"],
			LastName:     request.QueryStringParameters["lastname"],
			CreatedTime:  now.Format("01-02-2006 15:04:05"),
			ModifiedTime: now.Format("01-02-2006 15:04:05"),
		}

		insertId := middleware.InsertUser(userModel)
		res := PostResponse{
			ID:      insertId,
			Message: "User created successfully",
		}
		response, err := json.Marshal(&res)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}
		return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

	case "PATCH":
		fmt.Println("patch running")

		now := time.Now()
		id, intParseError := strconv.Atoi(request.QueryStringParameters["id"])

		if intParseError != nil {
			log.Fatalf("Unable to convert the string into int.  %v", intParseError)
		}
		userModel := models.User{
			FirstName:    request.QueryStringParameters["firstname"],
			LastName:     request.QueryStringParameters["lastname"],
			CreatedTime:  now.Format("01-02-2006 15:04:05"),
			ModifiedTime: now.Format("01-02-2006 15:04:05"),
		}

		updatedRows := middleware.UpdateUser(int64(id), userModel)

		msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
		res := PostResponse{
			ID:      int64(id),
			Message: msg,
		}
		response, err := json.Marshal(&res)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}
		return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

	case "DELETE":
		fmt.Println("delete running")

		id, intParseError := strconv.Atoi(request.QueryStringParameters["id"])

		if intParseError != nil {
			log.Fatalf("Unable to convert the string into int.  %v", intParseError)
		}
		deletedRows := middleware.DeleteUser(int64(id))

		msg := fmt.Sprintf("User Deleted%v", deletedRows)
		res := PostResponse{
			ID:      int64(id),
			Message: msg,
		}
		response, err := json.Marshal(&res)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}
		return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
	}
	return events.APIGatewayProxyResponse{Body: string("response"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
