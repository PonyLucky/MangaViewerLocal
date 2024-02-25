package manga_test

import (
	Manga "main/Manga"
	"testing"
)

func TestPage_GetPath(t *testing.T) {
	path := "/home/louis/Documents/Hakuneko/Mangas/The Other World’s Wizard Does Not Chant/Chapter 0001/01.jpg"
	page := Manga.Page{Path: path}
	expectedPath := path

	actualPath := page.GetPath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestPage_GetW(t *testing.T) {
	page := Manga.Page{W: 1200}
	expectedW := 1200

	actualW := page.GetW()

	if actualW != expectedW {
		t.Errorf("Expected width: %d, but got: %d", expectedW, actualW)
	}
}

func TestPage_GetH(t *testing.T) {
	page := Manga.Page{H: 800}
	expectedH := 800

	actualH := page.GetH()

	if actualH != expectedH {
		t.Errorf("Expected height: %d, but got: %d", expectedH, actualH)
	}
}

func TestPage_SetPath(t *testing.T) {
	page := Manga.Page{}
	expectedPath := "/home/louis/Documents/Hakuneko/Mangas/The Other World’s Wizard Does Not Chant/Chapter 0001/01.jpg"

	page.SetPath(expectedPath)

	if page.Path != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, page.Path)
	}
}

func TestPage_SetW(t *testing.T) {
	page := Manga.Page{}
	expectedW := 1200

	page.SetW(expectedW)

	if page.W != expectedW {
		t.Errorf("Expected width: %d, but got: %d", expectedW, page.W)
	}
}

func TestPage_SetH(t *testing.T) {
	page := Manga.Page{}
	expectedH := 800

	page.SetH(expectedH)

	if page.H != expectedH {
		t.Errorf("Expected height: %d, but got: %d", expectedH, page.H)
	}
}

func TestPage_LoadImageDimensions(t *testing.T) {
	page := Manga.Page{Path: "01.jpg"}
	basePath := "/home/louis/Documents/Hakuneko/Mangas/The Other World’s Wizard Does Not Chant/Chapter 0001"
	expectedW := 1200
	expectedH := 800

	page.LoadImageDimensions(basePath)

	actualW := page.GetW()
	actualH := page.GetH()

	if actualW != expectedW {
		t.Errorf("Expected width: %d, but got: %d", expectedW, actualW)
	}

	if actualH != expectedH {
		t.Errorf("Expected height: %d, but got: %d", expectedH, actualH)
	}
}
