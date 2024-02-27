package viewer

import (
	"encoding/json"
	"fmt"
	Manga "main/Manga"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Router is the main router for the viewer
type Router struct {
	Library *Manga.Library
}

func (r *Router) Serve() {
	http.HandleFunc("/", r.Handler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func (r *Router) Handler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	idxManga, idxChapter := -1, -1
	var err error

	fmt.Println("GET: '" + url + "'")

	parts := strings.Split(url, "/")
	for i := 0; i < len(parts); i++ {
		if parts[i] == "" {
			parts = append(parts[:i], parts[i+1:]...)
			i--
		}
	}

	if len(parts) > 0 {
		idxManga, err = strconv.Atoi(parts[0])
		if err != nil {
			idxManga = -1
		}
	}

	if len(parts) > 1 && err == nil {
		idxChapter, err = strconv.Atoi(parts[1])
		if err != nil {
			idxChapter = -1
		}
	}

	if idxManga == -1 {
		r.displayLibrary(w, req)
	} else if idxChapter == -1 {
		r.displayManga(w, req, idxManga)
	} else {
		r.displayChapter(w, req, idxManga, idxChapter)
	}
}

func (r *Router) displayLibrary(w http.ResponseWriter, req *http.Request) {
	// Convert the library to JSON
	jsonLibrary, _ := json.Marshal(r.Library)

	// Replace '%JSON%' in the file with the JSON
	file, err := os.ReadFile("./public/library/index.html")
	if err != nil {
		fmt.Println(err)
	}
	html := strings.Replace(string(file), "%JSON%", string(jsonLibrary), 1)

	// Write the HTML to the response
	w.Write([]byte(html))
}

func (r *Router) displayManga(w http.ResponseWriter, req *http.Request, index int) {
	manga := r.Library.GetManga(index)

	// Error handling if the manga does not exist
	if manga == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if manga.GetChapters() == nil {
		manga.LoadChapters(r.Library.GetPath())
	}

	jsonManga, _ := json.Marshal(manga)
	file, _ := os.ReadFile("./public/manga/index.html")
	html := strings.Replace(string(file), "%JSON%", string(jsonManga), 1)

	w.Write([]byte(html))
}

func (r *Router) displayChapter(w http.ResponseWriter, req *http.Request, mangaIndex int, chapterIndex int) {
	manga := r.Library.GetManga(mangaIndex)

	if manga.GetChapters() == nil {
		manga.LoadChapters(r.Library.GetPath())
	}

	chapter := manga.GetChapter(chapterIndex)

	// Error handling if the chapter does not exist
	if chapter == nil {
		http.Redirect(w, req, fmt.Sprintf("/%d", mangaIndex), http.StatusSeeOther)
		return
	}

	if chapter.GetPages() == nil {
		chapter.LoadPages(r.Library.GetPath())
	}

	jsonChapter, _ := json.Marshal(chapter)
	file, _ := os.ReadFile("./public/chapter/index.html")
	html := strings.Replace(string(file), "%JSON%", string(jsonChapter), 1)

	w.Write([]byte(html))
}
