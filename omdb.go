package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OMDBClient struct {
	APIKey  string
	BaseURL string
}

func NewOMDBClient(apiKey string) *OMDBClient {
	return &OMDBClient{
		APIKey:  apiKey,
		BaseURL: "http://www.omdbapi.com/",
	}
}

func (c *OMDBClient) GetMovie(imdbID string) (*Movie, error) {
	url := fmt.Sprintf("%s?apikey=%s&i=%s&plot=short", c.BaseURL, c.APIKey, imdbID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var omdbResp OMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&omdbResp); err != nil {
		return nil, err
	}

	if omdbResp.Response == "False" {
		return nil, fmt.Errorf("movie not found: %s", omdbResp.Error)
	}

	return &omdbResp.Movie, nil
}

var TopMovieIDs = []string{
	"tt0111161", // The Shawshank Redemption
	"tt0068646", // The Godfather
	"tt0071562", // The Godfather Part II
	"tt0468569", // The Dark Knight
	"tt0050083", // 12 Angry Men
	"tt0108052", // Schindler's List
	"tt0167260", // The Lord of the Rings: The Return of the King
	"tt0110912", // Pulp Fiction
	"tt0060196", // The Lord of the Rings: The Fellowship of the Ring
	"tt0137523", // Fight Club
	"tt0120737", // The Lord of the Rings: The Two Towers
	"tt0109830", // Forrest Gump
	"tt0080684", // Star Wars: The Empire Strikes Back
	"tt1375666", // Inception
	"tt0167261", // The Lord of the Rings: The Two Towers
	"tt0073486", // One Flew Over the Cuckoo's Nest
	"tt0099685", // Goodfellas
	"tt0047478", // Seven Samurai
	"tt0076759", // Star Wars
	"tt0317248", // City of God
	"tt0114369", // Se7en
	"tt0102926", // The Silence of the Lambs
	"tt0038650", // It's a Wonderful Life
	"tt0118799", // Life Is Beautiful
	"tt0114814", // The Usual Suspects
	"tt0245429", // Spirited Away
	"tt0120815", // Saving Private Ryan
	"tt0816692", // Interstellar
	"tt0110413", // Leon: The Professional
	"tt0120689", // The Green Mile
	"tt0317219", // Cars
	"tt0062622", // 2001: A Space Odyssey
	"tt0109831", // Forest Gump
	"tt0047396", // Rear Window
	"tt0082971", // Raiders of the Lost Ark
	"tt0078748", // Alien
	"tt0095327", // My Neighbor Totoro
	"tt0078788", // Apocalypse Now
	"tt0095765", // Cinema Paradiso
	"tt0172495", // Gladiator
	"tt0088763", // Back to the Future
	"tt0103064", // Terminator 2: Judgment Day
	"tt0482571", // The Prestige
	"tt0407887", // The Departed
	"tt0253474", // The Pianist
	"tt0027977", // Modern Times
	"tt0407154", // Casablanca
	"tt0209144", // Memento
	"tt0034583", // Citizen Kane
}
