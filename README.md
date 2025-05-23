# Movie Tinder 🎬

A simple "Tinder for movies" terminal app built with Go and BubbleTea. Swipe through top movies and keep track of what you like!

## Setup

1. Get a free OMDB API key from: http://www.omdbapi.com/apikey.aspx
2. Set your API key as an environment variable:
   ```bash
   export OMDB_API_KEY="your_api_key_here"
   ```

## Usage

```bash
# Build the app
go build -o movie-tinder .

# Run the app
./movie-tinder
```

## Controls

- **j** or **←** (left arrow): Dislike movie
- **l** or **→** (right arrow): Like movie  
- **v**: Toggle between movie swiping and viewing liked movies
- **q**: Quit the application

## Features

- Swipe through 50 top movies of all time
- Clean, minimal terminal interface
- Persistent storage of liked/disliked movies
- View your liked movies list
- Progress tracking

## Storage

Your movie preferences are saved in `movie_data.json` in the same directory.

## Dependencies

- [BubbleTea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- OMDB API - Movie data