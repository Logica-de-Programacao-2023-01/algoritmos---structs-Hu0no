package main

import (
	"fmt"
	"time"
)

type Song struct {
	Title    string
	Artist   string
	Duration time.Duration
}
type Playlist struct {
	Name  string
	Music []Song
}

func FindPlaylistsWithSong(playlists []Playlist, song Song) []string {
	PlaylistsWithSong := []string{}
	for _, playlist := range playlists {
		for _, music := range playlist.Music {
			if music == song {
				PlaylistsWithSong = append(PlaylistsWithSong, playlist.Name)
				break
			}
		}
	}
	return PlaylistsWithSong
}
func main() {
	playlist1 := Playlist{
		Name: "Rock Playlist",
		Music: []Song{
			{Title: "Song 1", Artist: "Artist 1", Duration: time.Minute},
			{Title: "Song 2", Artist: "Artist 2", Duration: 2 * time.Minute},
		},
	}
	playlist2 := Playlist{
		Name: "Pop Playlist",
		Music: []Song{
			{Title: "Song 3", Artist: "Artist 3", Duration: 3 * time.Minute},
			{Title: "Song 4", Artist: "Artist 4", Duration: 4 * time.Minute},
		},
	}
	playlist3 := Playlist{
		Name: "Heavy Metal Playlist",
		Music: []Song{
			{Title: "Song 2", Artist: "Artist 2", Duration: 2 * time.Minute},
			{Title: "Song 5", Artist: "Artist 5", Duration: 5 * time.Minute},
		},
	}
	SearchedSong := Song{
		Title:    "Song 2",
		Artist:   "Artist 2",
		Duration: 2 * time.Minute}
	playlists := []Playlist{playlist1, playlist2, playlist3}
	fmt.Println(FindPlaylistsWithSong(playlists, SearchedSong))
}
