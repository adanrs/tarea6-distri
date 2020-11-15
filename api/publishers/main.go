package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type BookRef struct {
	BookId int    `json:"book_id"`
	Title  string `json:"title"`
}

type Publisher struct {
	Id        int       `json:"_id"`
	Publisher string    `json:"publisher"`
	Country   string    `json:"country"`
	Founded   int       `json:"founded"`
	Genere    string    `json:"genere"`
	Books     []BookRef `json:"books"`
}

var items []Publisher

var jsonData string = `[
	{
		"_id": 1,
		"publisher": "John Wiley & Sons",
		"country": "United States",
		"founded": 1807,
		"genere": "Academic",
		"books": [
			{
				"book_id": 1,
				"title": "Operating System Concepts"
			},
			{
				"book_id": 2,
				"title": "Database System Concepts"
			}
		]
	},
	{
		"_id": 2,
		"publisher": "Pearson Education",
		"country": "United Kingdom",
		"founded": 1844,
		"genere": "Education",
		"books": [
			{
				"book_id": 3,
				"title": "Computer Networks"
			},
			{
				"book_id": 4,
				"title": "Modern Operating Systems"
			}
		]
	}
]`

func FindItem(id int) *Publisher {
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
