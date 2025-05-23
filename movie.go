package main

type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Plot       string `json:"Plot"`
	Director   string `json:"Director"`
	Actors     string `json:"Actors"`
	Genre      string `json:"Genre"`
	ImdbRating string `json:"imdbRating"`
}

type OMDBResponse struct {
	Movie
	Response string `json:"Response"`
	Error    string `json:"Error"`
}
