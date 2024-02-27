package manga

import (
	"os"
)

/*
Representation of a Library. (A library is a collection of mangas)

# Fields
  - Path:      The path to the library folder.
  - Mangas:    The list of mangas in the library.
*/
type Library struct {
	Path   string
	Mangas []Manga
}

// ------------------------------
// Getters and Setters
// ------------------------------

/*
Get the path to the library folder.

# Returns
  - The path to the library folder.
*/
func (l *Library) GetPath() string {
	return l.Path
}

/*
Get the list of mangas in the library.

# Returns
  - The list of mangas in the library.
*/
func (l *Library) GetMangas() *[]Manga {
	return &l.Mangas
}

/*
Set the path to the library folder.

# Params
- path: The path to the library folder.
*/
func (l *Library) SetPath(path string) {
	l.Path = path
}

/*
Set the list of mangas in the library.

# Params
- mangas: The list of mangas in the library.
*/
func (l *Library) SetMangas(mangas *[]Manga) {
	l.Mangas = *mangas
}

// ------------------------------
// Methods
// ------------------------------

/*
Get specific manga from the library.

# Params
- index: The index of the manga in the library.

# Returns
  - The manga at the given index. If the index is out of range, it will return nil.
*/
func (l *Library) GetManga(index int) *Manga {
	if index < 0 || index >= len(l.Mangas) {
		return nil
	}
	return &(*l.GetMangas())[index]
}

/*
Init the library from the given path.

# Params
- path: The path to the library folder.
*/
func (l *Library) Init() {
	mangas := []Manga{}

	files, _ := os.ReadDir(l.Path)
	for _, file := range files {
		if file.IsDir() {
			manga := Manga{Path: file.Name()}
			manga.Init(l.Path)
			mangas = append(mangas, manga)
		}
	}

	l.SetMangas(&mangas)
}
