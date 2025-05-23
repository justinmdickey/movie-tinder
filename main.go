package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	omdbClient     *OMDBClient
	storage        *Storage
	currentMovie   *Movie
	movieIndex     int
	loading        bool
	error          string
	unseenMovies   []string
	showLikedList  bool
	likedMovies    []*Movie
	width          int
	height         int
}

var (
	borderStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("211")).
		Bold(true).
		Align(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("243"))

	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Italic(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true)

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("46")).
		Bold(true)

	superlikeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("201")).
		Bold(true)
)

type movieFetched struct {
	movie *Movie
	err   error
}

func initialModel(apiKey string) model {
	storage := LoadStorage()
	
	var unseenMovies []string
	for _, id := range TopMovieIDs {
		if !storage.IsSeen(id) {
			unseenMovies = append(unseenMovies, id)
		}
	}

	return model{
		omdbClient:   NewOMDBClient(apiKey),
		storage:      storage,
		loading:      true,
		unseenMovies: unseenMovies,
		movieIndex:   0,
	}
}

func (m model) Init() tea.Cmd {
	if len(m.unseenMovies) > 0 {
		return fetchMovie(m.omdbClient, m.unseenMovies[0])
	}
	return nil
}

func fetchMovie(client *OMDBClient, imdbID string) tea.Cmd {
	return func() tea.Msg {
		movie, err := client.GetMovie(imdbID)
		return movieFetched{movie: movie, err: err}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "j":
			if m.currentMovie != nil && !m.loading {
				return m.dislikeMovie()
			}
		case "right", "l":
			if m.currentMovie != nil && !m.loading {
				return m.likeMovie()
			}
		case "up", "k":
			if m.currentMovie != nil && !m.loading {
				return m.superlikeMovie()
			}
		case "down", "d":
			if m.currentMovie != nil && !m.loading {
				return m.markAsNotSeen()
			}
		case "v":
			return m.toggleLikedList()
		}

	case movieFetched:
		if msg.err != nil {
			m.error = msg.err.Error()
			m.loading = false
		} else {
			m.currentMovie = msg.movie
			m.loading = false
			m.error = ""
		}
		return m, nil
	}

	return m, nil
}

func (m model) likeMovie() (model, tea.Cmd) {
	m.storage.LikeMovie(m.currentMovie.ImdbID)
	m.storage.SaveStorage()
	return m.nextMovie()
}

func (m model) dislikeMovie() (model, tea.Cmd) {
	m.storage.DislikeMovie(m.currentMovie.ImdbID)
	m.storage.SaveStorage()
	return m.nextMovie()
}

func (m model) superlikeMovie() (model, tea.Cmd) {
	m.storage.SuperlikeMovie(m.currentMovie.ImdbID)
	m.storage.SaveStorage()
	return m.nextMovie()
}

func (m model) markAsNotSeen() (model, tea.Cmd) {
	m.storage.MarkAsNotSeen(m.currentMovie.ImdbID)
	m.storage.SaveStorage()
	return m.nextMovie()
}

func (m model) nextMovie() (model, tea.Cmd) {
	m.movieIndex++
	if m.movieIndex >= len(m.unseenMovies) {
		m.currentMovie = nil
		return m, nil
	}
	
	m.loading = true
	return m, fetchMovie(m.omdbClient, m.unseenMovies[m.movieIndex])
}

func (m model) toggleLikedList() (model, tea.Cmd) {
	if m.showLikedList {
		m.showLikedList = false
		return m, nil
	}

	m.showLikedList = true
	var likedMovies []*Movie
	
	for _, imdbID := range m.storage.SuperlikedMovies {
		movie, err := m.omdbClient.GetMovie(imdbID)
		if err == nil {
			likedMovies = append(likedMovies, movie)
		}
	}
	
	for _, imdbID := range m.storage.LikedMovies {
		movie, err := m.omdbClient.GetMovie(imdbID)
		if err == nil {
			likedMovies = append(likedMovies, movie)
		}
	}
	
	m.likedMovies = likedMovies
	return m, nil
}

func (m model) View() string {
	if m.showLikedList {
		return m.renderLikedList()
	}

	if m.loading {
		content := "Loading movie..."
		return m.wrapContent(content)
	}

	if m.error != "" {
		content := errorStyle.Render("Error: "+m.error) + "\n\n" + helpStyle.Render("Press 'q' to quit")
		return m.wrapContent(content)
	}

	if m.currentMovie == nil {
		content := successStyle.Render("All movies completed!") + "\n\n" + 
			helpStyle.Render("Press 'v' to view liked movies\nPress 'q' to quit")
		return m.wrapContent(content)
	}

	return m.renderMovie()
}

func (m model) renderMovie() string {
	movie := m.currentMovie
	
	header := headerStyle.Render("MOVIE TINDER")
	title := titleStyle.Render(fmt.Sprintf("%s (%s)", movie.Title, movie.Year))
	
	info := fmt.Sprintf("Rating: %s  |  %s", movie.ImdbRating, movie.Genre)
	details := fmt.Sprintf("Director: %s\nCast: %s", movie.Director, movie.Actors)
	
	plot := movie.Plot
	if len(plot) > 250 {
		plot = plot[:250] + "..."
	}
	
	controls := "[j/←] Dislike  [l/→] Like  [k/↑] Superlike  [d/↓] Not Seen  [v] View Liked  [q] Quit"
	progress := fmt.Sprintf("Progress: %d/%d", m.movieIndex+1, len(TopMovieIDs))
	
	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		title,
		subtitleStyle.Render(info),
		"",
		subtitleStyle.Render(details),
		"",
		plot,
		"",
		helpStyle.Render(controls),
		"",
		subtitleStyle.Render(progress),
	)
	
	return m.wrapContent(content)
}

func (m model) renderLikedList() string {
	header := headerStyle.Render("LIKED MOVIES")
	
	var movieList strings.Builder
	if len(m.likedMovies) == 0 {
		movieList.WriteString(subtitleStyle.Render("No liked movies yet!"))
	} else {
		count := 1
		
		if len(m.storage.SuperlikedMovies) > 0 {
			movieList.WriteString(superlikeStyle.Render("SUPERLIKED:") + "\n")
			for _, imdbID := range m.storage.SuperlikedMovies {
				for _, movie := range m.likedMovies {
					if movie.ImdbID == imdbID {
						entry := fmt.Sprintf("%d. %s (%s) - %s ★", 
							count, movie.Title, movie.Year, movie.ImdbRating)
						movieList.WriteString(superlikeStyle.Render(entry) + "\n")
						count++
						break
					}
				}
			}
			movieList.WriteString("\n")
		}
		
		if len(m.storage.LikedMovies) > 0 {
			movieList.WriteString(titleStyle.Render("LIKED:") + "\n")
			for _, imdbID := range m.storage.LikedMovies {
				for _, movie := range m.likedMovies {
					if movie.ImdbID == imdbID {
						entry := fmt.Sprintf("%d. %s (%s) - %s", 
							count, movie.Title, movie.Year, movie.ImdbRating)
						movieList.WriteString(entry + "\n")
						count++
						break
					}
				}
			}
		}
	}
	
	controls := "[v] Back to swiping  [q] Quit"
	
	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		movieList.String(),
		"",
		helpStyle.Render(controls),
	)
	
	return m.wrapContent(content)
}

func (m model) wrapContent(content string) string {
	if m.width == 0 {
		m.width = 80
	}
	
	contentWidth := min(m.width-4, 80)
	style := borderStyle.Width(contentWidth)
	
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, style.Render(content))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set OMDB_API_KEY environment variable")
		fmt.Println("Get your free API key at: http://www.omdbapi.com/apikey.aspx")
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(apiKey), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}