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
	"tt0062622", // 2001: A Space Odyssey
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
	"tt0209144", // Memento
	"tt0034583", // Citizen Kane
	"tt0042192", // Casablanca
	"tt0095016", // Die Hard
	"tt0087843", // Once Upon a Time in the West
	"tt0057012", // Dr. Strangelove
	"tt0033467", // Citizen Kane
	"tt0090605", // Aliens
	"tt0053125", // North by Northwest
	"tt0051201", // Vertigo
	"tt0086190", // Star Wars: Return of the Jedi
	"tt0075314", // Taxi Driver
	"tt0086879", // Amadeus
	"tt0112573", // Braveheart
	"tt0119698", // Saving Private Ryan
	"tt0119217", // Good Will Hunting
	"tt0105236", // Reservoir Dogs
	"tt0087544", // The Terminator
	"tt0361748", // Inglourious Basterds
	"tt0910970", // WALL-E
	"tt0043014", // Sunset Blvd.
	"tt0056058", // Lawrence of Arabia
	"tt0364569", // Oldboy
	"tt0993846", // The Wolf of Wall Street
	"tt0081505", // The Shining
	"tt0071853", // The Exorcist
	"tt0053604", // Psycho
	"tt0032553", // The Great Dictator
	"tt0405094", // Lives of Others
	"tt0052357", // Singin' in the Rain
	"tt0066921", // A Clockwork Orange
	"tt0044741", // Singing in the Rain
	"tt0036775", // Double Indemnity
	"tt0045152", // Singin' in the Rain
	"tt0986264", // Toy Story 3
	"tt0211915", // Amélie
	"tt0093058", // Full Metal Jacket
	"tt0097576", // Indiana Jones and the Last Crusade
	"tt0015864", // The Gold Rush
	"tt0064116", // Once Upon a Time in the West
	"tt0167404", // The Sixth Sense
	"tt0338013", // Eternal Sunshine of the Spotless Mind
	"tt0101414", // Beauty and the Beast
	"tt0022100", // M
	"tt0091251", // Ferris Bueller's Day Off
	"tt0077416", // Annie Hall
	"tt0169547", // American Beauty
	"tt0208092", // Snatch
	"tt0105695", // Silence of the Lambs
	"tt0096283", // My Neighbor Totoro
	"tt0036868", // The Third Man
	"tt0092005", // The Princess Bride
	"tt0086250", // Scarface
	"tt0070047", // The Sting
	"tt0113277", // Heat
	"tt0758758", // Into the Wild
	"tt0033870", // The Maltese Falcon
	"tt0044079", // Strangers on a Train
	"tt0363163", // The Incredibles
	"tt0077711", // Stalker
	"tt0083658", // Blade Runner
	"tt0180093", // Requiem for a Dream
	"tt0112641", // Casino
	"tt0088247", // Once Upon a Time in America
}
