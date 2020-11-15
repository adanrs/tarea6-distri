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

type Actors struct {
	Id     int        `json:"_id"`
	Actors string     `json:"actors"`
	Movies []MovieRef `json:"movies"`
}

var items []Actors

var jsonData string = `[
	{
		"_id": 1,
		"Actors": "Chris Pratt, Vin Diesel, Bradley Cooper, Zoe Saldana",
		"movies": [
			{
				"movie_id": 1,
				"title": "Guardians of the Galaxy"
			}
		]
	}
]`

func FindItem(id int) *Actors {
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
