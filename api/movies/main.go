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
	},
	{
		"_id": 2,
		"title": "Prometheus",
		"Genre": "Adventure,Mystery,Sci-Fi",
		"Description": "Following clues to the origin of mankind, a team finds a structure on a distant moon, but they soon realize they are not alone.",
		"language": "ENGLISH",
		"Director": "Ridley Scott",
		"Actors": "Noomi Rapace, Logan Marshall-Green, Michael Fassbender, Charlize Theron",
		"Director_id": 2,
		"Actors_id": 2,
		"Year": 2012
	},
	{
		"_id": 3,
		"title": "Split",
		"Genre": "Horror,Thriller",
		"Description": "Three girls are kidnapped by a man with a diagnosed 23 distinct personalities. They must try to escape before the apparent emergence of a frightful new 24th.",
		"language": "ENGLISH",
		"Director": "M. Night Shyamalan",
		"Actors": "James McAvoy, Anya Taylor-Joy, Haley Lu Richardson, Jessica Sula",
		"Director_id": 3,
		"Actors_id": 3,
		"Year": 2016
	}
]`

func FindMovie(id int) *Movie {
	for _, movie := range movies {
		if movie.Id == id {
			return &movie
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
			movie := FindMovie(param)
			if movie != nil {
				data, _ = json.Marshal(*movie)
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
