package manga

import (
	Misc "main/Misc"
	"os"
	"path/filepath"
)

/*
Representation of a manga.

# Fields
  - Path:      The path to the manga folder.
  - Title:     The title of the manga.
  - Poster:    The poster of the manga. If not available, it will be nil.
  - Chapters:  The list of chapters in the manga.
*/
type Manga struct {
	Path     string
	Title    string
	Poster   string
	Chapters []Chapter
}

// ------------------------------
// Getters and Setters
// ------------------------------

/*
Get the path to the manga folder.

# Returns
  - The path to the manga folder.
*/
func (m *Manga) GetPath() string {
	return m.Path
}

/*
Get the title of the manga.

# Returns
  - The title of the manga.
*/
func (m *Manga) GetTitle() *string {
	return &m.Title
}

/*
Get the list of chapters in the manga.

# Returns
  - The list of chapters in the manga.
*/
func (m *Manga) GetChapters() *[]Chapter {
	return &m.Chapters
}

/*
Get the poster of the manga.

# Return
  - The poster of the manga. If not available, it will be nil.
*/
func (m *Manga) GetPoster() *string {
	return &m.Poster
}

/*
Set the path to the manga folder.

# Params
  - path: The path to the manga folder.
*/
func (m *Manga) SetPath(path string) {
	m.Path = path
}

/*
Set the title of the manga.

# Params
  - title: The title of the manga.

# Note
  - The title will be formatted using the Misc.FormatTitle function.
*/
func (m *Manga) SetTitle(title string) {
	m.Title = Misc.FormatTitle(title)
}

/*
Set the list of chapters in the manga.

# Params
  - chapters: The list of chapters in the manga.
*/
func (m *Manga) SetChapters(chapters *[]Chapter) {
	m.Chapters = *chapters
}

/*
Set the poster of the manga.

# Params
  - poster: The poster of the manga.
*/
func (m *Manga) SetPoster(poster *string) {
	m.Poster = *poster
}

// ------------------------------
// Methods
// ------------------------------

/*
Load all chapters in the manga.

# Params
  - basePath: The base path to the manga folder.
*/
func (m *Manga) LoadChapters(basePath string) {
	var chapters []Chapter

	files, _ := os.ReadDir(filepath.Join(basePath, m.GetPath()))
	for _, file := range files {
		if file.IsDir() {
			chapter := Chapter{Path: file.Name()}
			chapter.Init()
			chapters = append(chapters, chapter)
		}
	}

	m.SetChapters(&chapters)
}

/*
Load the poster of the manga.

# Params
  - basePath: The base path to the manga folder.
*/
func (m *Manga) LoadPoster(basePath string) {
	exts := []string{"jpg", "jpeg", "png", "webp"}
	var poster string

	for _, ext := range exts {
		name := "poster." + ext
		path := filepath.Join(basePath, m.GetPath(), name)
		if _, err := os.Stat(path); err == nil {
			poster = name
			break
		}
	}

	m.SetPoster(&poster)
}

/*
Get the chapter at the specified index.

# Params
  - index: The index of the chapter.

# Returns
  - The chapter at the specified index. If the index is out of range,
    it will return an empty chapter.
*/
func (m *Manga) GetChapter(index int) *Chapter {
	if index < 0 || index >= m.GetNumberOfChapters() {
		return &Chapter{}
	}
	return &(*m.GetChapters())[index]
}

/*
Get the number of chapters in the manga.

# Returns
  - The number of chapters in the manga.
*/
func (m *Manga) GetNumberOfChapters() int {
	return len(*m.GetChapters())
}

/*
Initialize the manga.

# Params
  - basePath: The base path to the manga folder.
*/
func (m *Manga) Init(basePath string) {
	m.SetTitle(filepath.Base(m.GetPath()))
	m.LoadPoster(basePath)
}
