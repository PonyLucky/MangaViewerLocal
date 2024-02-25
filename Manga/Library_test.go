package manga_test

import (
	Manga "main/Manga"
	"testing"
)

func TestLibrary_GetPath(t *testing.T) {
	path := "/home/louis/Documents/Hakuneko/Mangas"
	library := Manga.Library{Path: path}
	expectedPath := path

	actualPath := library.GetPath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestLibrary_GetMangas(t *testing.T) {
	manga1 := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	manga2 := Manga.Manga{Path: "THE REINCARNATION MAGICIAN OF THE INFERIOR EYES"}
	library := Manga.Library{Mangas: []Manga.Manga{manga1, manga2}}
	expectedMangas := []Manga.Manga{manga1, manga2}

	actualMangas := library.GetMangas()

	if len(*actualMangas) != len(expectedMangas) {
		t.Errorf("Expected %d mangas, but got %d", len(expectedMangas), len(*actualMangas))
	}

	for i, manga := range expectedMangas {
		if (*actualMangas)[i].Path != manga.Path {
			t.Errorf("Expected manga at index %d: %s, but got: %s", i, manga.Path, (*actualMangas)[i].Path)
		}
	}
}

func TestLibrary_SetPath(t *testing.T) {
	path := "/home/louis/Documents/Hakuneko/Mangas"
	library := Manga.Library{}
	expectedPath := path

	library.SetPath(path)

	actualPath := library.GetPath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestLibrary_SetMangas(t *testing.T) {
	manga1 := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	manga2 := Manga.Manga{Path: "THE REINCARNATION MAGICIAN OF THE INFERIOR EYES"}
	library := Manga.Library{}
	expectedMangas := []Manga.Manga{manga1, manga2}

	library.SetMangas(&expectedMangas)

	actualMangas := library.GetMangas()

	if len(*actualMangas) != len(expectedMangas) {
		t.Errorf("Expected %d mangas, but got %d", len(expectedMangas), len(*actualMangas))
	}

	for i, manga := range expectedMangas {
		if (*actualMangas)[i].Path != manga.Path {
			t.Errorf("Expected manga at index %d: %s, but got: %s", i, manga.Path, (*actualMangas)[i].Path)
		}
	}
}

func TestLibrary_GetManga(t *testing.T) {
	manga1 := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	manga2 := Manga.Manga{Path: "THE REINCARNATION MAGICIAN OF THE INFERIOR EYES"}
	library := Manga.Library{Mangas: []Manga.Manga{manga1, manga2}}
	expectedManga := manga1

	actualManga := library.GetManga(0)

	if actualManga.Path != expectedManga.Path {
		t.Errorf("Expected manga: %s, but got: %s", expectedManga.Path, actualManga.Path)
	}
}

func TestLibrary_Init(t *testing.T) {
	path := "/home/louis/Documents/Hakuneko/Mangas"
	library := Manga.Library{Path: path}
	manga1 := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	manga2 := Manga.Manga{Path: "THE REINCARNATION MAGICIAN OF THE INFERIOR EYES"}
	expectedMangas := []Manga.Manga{manga2, manga1}

	library.Init()

	actualMangas := library.GetMangas()

	if len(*actualMangas) != len(expectedMangas) {
		t.Errorf("Expected %d mangas, but got %d", len(expectedMangas), len(*actualMangas))
	}

	for i, manga := range expectedMangas {
		if (*actualMangas)[i].Path != manga.Path {
			t.Errorf("Expected manga at index %d: %s, but got: %s", i, manga.Path, (*actualMangas)[i].Path)
		}
	}
}
