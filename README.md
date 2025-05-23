# Movie Tinder 🎬

Test

A terminal-based "Tinder for movies" app built with Go and Bubble Tea. Discover and rate top movies with an intuitive swipe-like interface, all from your command line!

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
   git clone https://github.com/yourusername/movie-tinder.git
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
- **r**: Reset all movie preferences and start over
- **q**: Quit the application

## Features

- 🎭 Discover 100 top movies of all time
- ✨ Beautiful terminal interface with smooth styling
- 💾 Persistent storage of your movie preferences
- ⭐ Superlike system for absolute favorites
- 🔄 "Not seen" option to re-encounter movies later
- 📋 View your curated list of liked movies by category
- 🔄 Reset functionality to start fresh anytime
- 📊 Progress tracking through the movie collection
- 🚀 Cross-platform support (Linux, macOS, Windows)

## Storage

Your movie preferences are saved in `movie_data.json` in the same directory.

## Technologies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Modern terminal UI framework for Go
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Styling and layout library
- **[OMDB API](http://www.omdbapi.com/)** - Movie database for rich film information

## Contributing

Feel free to open issues or submit pull requests! This project was built as a fun exploration of terminal UI development with Go.

## License

MIT License - feel free to use this project as inspiration for your own terminal applications!

