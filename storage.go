package main

import (
	"encoding/json"
	"os"
)

type Storage struct {
	LikedMovies      []string `json:"liked_movies"`
	SuperlikedMovies []string `json:"superliked_movies"`
	SeenMovies       []string `json:"seen_movies"`
}

const storageFile = "movie_data.json"

func LoadStorage() *Storage {
	storage := &Storage{
		LikedMovies:      []string{},
		SuperlikedMovies: []string{},
		SeenMovies:       []string{},
	}

	data, err := os.ReadFile(storageFile)
	if err != nil {
		return storage
	}

	json.Unmarshal(data, storage)
	return storage
}

func (s *Storage) SaveStorage() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func (s *Storage) LikeMovie(imdbID string) {
	if !s.IsLiked(imdbID) && !s.IsSuperliked(imdbID) {
		s.LikedMovies = append(s.LikedMovies, imdbID)
	}
	s.MarkAsSeen(imdbID)
}

func (s *Storage) SuperlikeMovie(imdbID string) {
	s.removeLike(imdbID)
	if !s.IsSuperliked(imdbID) {
		s.SuperlikedMovies = append(s.SuperlikedMovies, imdbID)
	}
	s.MarkAsSeen(imdbID)
}

func (s *Storage) MarkAsNotSeen(imdbID string) {
	s.removeFromSeen(imdbID)
	s.removeLike(imdbID)
	s.removeSuperlike(imdbID)
}

func (s *Storage) DislikeMovie(imdbID string) {
	s.MarkAsSeen(imdbID)
}

func (s *Storage) MarkAsSeen(imdbID string) {
	for _, id := range s.SeenMovies {
		if id == imdbID {
			return
		}
	}
	s.SeenMovies = append(s.SeenMovies, imdbID)
}

func (s *Storage) IsLiked(imdbID string) bool {
	for _, id := range s.LikedMovies {
		if id == imdbID {
			return true
		}
	}
	return false
}

func (s *Storage) IsSuperliked(imdbID string) bool {
	for _, id := range s.SuperlikedMovies {
		if id == imdbID {
			return true
		}
	}
	return false
}

func (s *Storage) removeLike(imdbID string) {
	for i, id := range s.LikedMovies {
		if id == imdbID {
			s.LikedMovies = append(s.LikedMovies[:i], s.LikedMovies[i+1:]...)
			return
		}
	}
}

func (s *Storage) removeSuperlike(imdbID string) {
	for i, id := range s.SuperlikedMovies {
		if id == imdbID {
			s.SuperlikedMovies = append(s.SuperlikedMovies[:i], s.SuperlikedMovies[i+1:]...)
			return
		}
	}
}

func (s *Storage) removeFromSeen(imdbID string) {
	for i, id := range s.SeenMovies {
		if id == imdbID {
			s.SeenMovies = append(s.SeenMovies[:i], s.SeenMovies[i+1:]...)
			return
		}
	}
}

func (s *Storage) IsSeen(imdbID string) bool {
	for _, id := range s.SeenMovies {
		if id == imdbID {
			return true
		}
	}
	return false
}
