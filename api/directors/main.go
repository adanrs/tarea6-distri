package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MovieRef struct {
	MovieId int    `json:"movie_id"`
	Title   string `json:"title"`
}

type Directors struct {
	Id       int    `json:"_id"`
	Director string `json:"director"`

	Movie []MovieRef `json:"movies"`
}

var items []Directors

var jsonData string = `[
	{
		"_id": 1,
		"director": "James Gunn",
		"movies": [
			{
				"movie_id": 1,
				"title": "Guardians of the Galaxy"
			}
		]
	},
	{
		"_id": 2,
"director": "Ridley Scott",

		"movies": [
			{
				"movie_id": 2,
				"title": "Prometheus"
			}
		]
	},
	{
		"_id": 3,
		"director": "M. Night Shyamalan",
	"movies": [
		{
			"movie_id": 3,
			"title": "Split"
		}
	]
	}
]`

func FindItem(id int) *Directors {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
