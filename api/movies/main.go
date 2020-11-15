package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	Id          int    `json:"_id"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Description string `json:"description"`
	Genre       int    `json:"genre"`
	Language    string `json:"language"`
	Actors      string `json:"actors"`
	Actors_Id   int    `json:"actors_id"`
	Director    string `json:"director"`
	Director_Id int    `json:"director_id"`
}

var movies []Movie

var jsonData string = `[
	{
		"_id": 1,
		"title": "Guardians of the Galaxy",
		"Genre": "Action,Adventure,Sci-Fi",
		"Description": "A group of intergalactic criminals are forced to work together to stop a fanatical warrior from taking control of the universe.",
		"language": "ENGLISH",
		"Director": "James Gunn",
		"Actors": "Chris Pratt, Vin Diesel, Bradley Cooper, Zoe Saldana",
		"Director_id": 1,
		"Actors_id": 1,
		"Year": 2014	
	}
]`

func FindBook(id int) *Movie {
	for _, book := range movies {
		if book.Id == id {
			return &book
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(movies)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			book := FindBook(param)
			if book != nil {
				data, _ = json.Marshal(*book)
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
	_ = json.Unmarshal([]byte(jsonData), &movies)
	lambda.Start(handler)
}
