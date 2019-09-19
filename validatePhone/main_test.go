package main

import (
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func mapToString(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k+"="+m[k])
	}
	return "[" + strings.Join(keys, ", ") + "]"
}

func TestHandler(t *testing.T) {
	saved := getPhoneValidation
	defer func() { getPhoneValidation = saved }()

	getPhoneValidation = func(phoneNumber string) (response *APIResponse, err error) {
		response = &APIResponse{
			Valid:    true,
			Number:   phoneNumber,
			LineType: "mobile",
		}

		return response, nil
	}

	tests := []struct {
		test string
		events.APIGatewayProxyRequest
		want int
	}{
		{
			"should respond successfully (200) when a valid phoneNumber is provided",
			events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{"phoneNumber": "17871234567"}},
			200,
		},
		{
			"should respond with bad request (400) when no phone number is provided",
			events.APIGatewayProxyRequest{},
			400,
		},
	}

	for _, test := range tests {
		response, _ := Handler(test.APIGatewayProxyRequest)

		qt := mapToString(test.QueryStringParameters)

		if got := response.StatusCode; got != test.want {
			t.Errorf("Handler() = %d, want %d\t querystring: %s", got, test.want, qt)
		}
	}
}
