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

func PlaylistInfo(playlist Playlist) int {
	fmt.Println("Playlist Name: ", playlist.Name)
	total_time := 0 * time.Second
	for _, song := range playlist.Music {
		fmt.Println("Song Title: ", song.Title)
		fmt.Println("Song Artist: ", song.Artist)
		fmt.Println("Song Duration: ", song.Duration)
		total_time += song.Duration
	}
	fmt.Println("Playlist total time: ", total_time)
	return 0
}
func main() {
	playlist := Playlist{
		Name: "Rock",
		Music: []Song{
			{Title: "505", Artist: "Arctic Monkeys", Duration: 4*time.Minute + 13*time.Second},
			{Title: "Heaven", Artist: "Bryan Adams", Duration: 4*time.Minute + 11*time.Second},
			{Title: "Monster", Artist: "Skillet", Duration: 3*time.Minute + 6*time.Second},
		},
	}
	fmt.Print(PlaylistInfo(playlist))
}
