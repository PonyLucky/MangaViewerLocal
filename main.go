package main

import (
	// "encoding/json"
	// "path/filepath"
	Manga "main/Manga"
	Viewer "main/Viewer"
)

func main() {
	// l := Manga.Library{Path: "/home/louis/Documents/Hakuneko/Mangas"}
	// l.Init()

	// // Get first manga and load its chapters
	// manga := l.GetManga(0)
	// manga.LoadChapters(l.GetPath())
	// mangaPath := filepath.Join(l.GetPath(), manga.GetPath())

	// // Get first chapter and load its pages
	// chapter := manga.GetChapter(0)
	// chapter.LoadPages(mangaPath)

	// // Convert the library to JSON
	// json, _ := json.MarshalIndent(l, "", "  ")

	// // Print the JSON
	// println(string(json))

	l := Manga.Library{Path: "/home/louis/Documents/Hakuneko/Mangas"}
	l.Init()

	v := Viewer.Router{Library: &l}
	v.Serve()
}
