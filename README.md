# Movie Tinder 🎬

A simple "Tinder for movies" terminal app built with Go and BubbleTea. Swipe through top movies and keep track of what you like!

## Installation

### Option 1: Download Pre-built Binary (Recommended)

1. Get a free OMDB API key from: http://www.omdbapi.com/apikey.aspx
2. Download the latest release for your platform from [Releases](../../releases)
3. Extract and run:
   ```bash
   # Set your API key
   export OMDB_API_KEY="your_api_key_here"
   
   # Run the app
   ./movie-tinder-*
   ```

### Option 2: Build from Source

1. Get a free OMDB API key from: http://www.omdbapi.com/apikey.aspx
2. Clone and build:
   ```bash
   git clone <repo-url>
   cd movie-tinder
   go build -o movie-tinder .
   export OMDB_API_KEY="your_api_key_here"
   ./movie-tinder
   ```

## Controls

- **j** or **←** (left arrow): Dislike movie
- **l** or **→** (right arrow): Like movie
- **k** or **↑** (up arrow): Superlike movie (highlighted in pink)
- **d** or **↓** (down arrow): Mark as not seen (encounter again)
- **v**: Toggle between movie swiping and viewing liked movies
- **q**: Quit the application

## Features

- Swipe through 50 top movies of all time
- Clean, minimal terminal interface with lipgloss styling
- Persistent storage of liked/superliked/disliked movies
- Superlike system for favorite movies
- "Not seen" option to re-encounter movies
- View your liked movies list with categories
- Progress tracking

## Storage

Your movie preferences are saved in `movie_data.json` in the same directory.

## Dependencies

- [BubbleTea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
- OMDB API - Movie data