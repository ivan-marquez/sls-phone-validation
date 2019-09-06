package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	phoneNumber, ok := request.QueryStringParameters["phoneNumber"]

	if !ok {
		return Response{
			StatusCode: 400,
			Body:       "phoneNumber not provided",
		}, nil
	}

	res, err := ValidateMobilePhone(phoneNumber)
	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	body, err := json.Marshal(res)

	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	json.HTMLEscape(&buf, body)

	headers := map[string]string{
		"Content-Type":           "application/json",
		"X-MyCompany-Func-Reply": "hello-handler",
	}

	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers:         headers,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
