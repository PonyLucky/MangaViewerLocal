package manga_test

import (
	Manga "main/Manga"
	"testing"
)

func TestChapter_GetPath(t *testing.T) {
	path := "Chapter 0001"
	chapter := Manga.Chapter{Path: path}
	expectedPath := path

	actualPath := chapter.GetPath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestChapter_GetTitle(t *testing.T) {
	chapter := Manga.Chapter{Title: "Chapter 1"}
	expectedTitle := "Chapter 1"

	actualTitle := chapter.GetTitle()

	if actualTitle != expectedTitle {
		t.Errorf("Expected title: %s, but got: %s", expectedTitle, actualTitle)
	}
}

func TestChapter_GetPages(t *testing.T) {
	page1 := Manga.Page{Path: "01.jpg"}
	page2 := Manga.Page{Path: "02.jpg"}
	chapter := Manga.Chapter{Pages: []Manga.Page{page1, page2}}
	expectedPages := []Manga.Page{page1, page2}

	actualPages := chapter.GetPages()

	if len(*actualPages) != len(expectedPages) {
		t.Errorf("Expected %d pages, but got %d", len(expectedPages), len(*actualPages))
	}

	for i, page := range expectedPages {
		if (*actualPages)[i] != page {
			t.Errorf("Expected page at index %d: %s, but got: %s", i, page.Path, (*actualPages)[i].Path)
		}
	}
}

func TestChapter_SetPath(t *testing.T) {
	chapter := Manga.Chapter{}
	expectedPath := "Chapter 0001"

	chapter.SetPath(expectedPath)

	if chapter.Path != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, chapter.Path)
	}
}

func TestChapter_SetTitle(t *testing.T) {
	chapter := Manga.Chapter{}
	expectedTitle := "Chapter 1"

	chapter.SetTitle(expectedTitle)

	if chapter.Title != expectedTitle {
		t.Errorf("Expected title: %s, but got: %s", expectedTitle, chapter.Title)
	}
}

func TestChapter_SetPages(t *testing.T) {
	page1 := Manga.Page{Path: "01.jpg"}
	page2 := Manga.Page{Path: "02.jpg"}
	chapter := Manga.Chapter{}
	expectedPages := []Manga.Page{page1, page2}

	chapter.SetPages(&expectedPages)

	if len(chapter.Pages) != len(expectedPages) {
		t.Errorf("Expected %d pages, but got %d", len(expectedPages), len(chapter.Pages))
	}

	for i, page := range expectedPages {
		if chapter.Pages[i] != page {
			t.Errorf("Expected page at index %d: %s, but got: %s", i, page.Path, chapter.Pages[i].Path)
		}
	}
}

func TestChapter_LoadPages(t *testing.T) {
	path := "Chapter 0001"
	chapter := Manga.Chapter{Path: path}
	chapter.LoadPages("/home/louis/Documents/Hakuneko/Mangas/The Other World’s Wizard Does Not Chant/")
	if len(chapter.Pages) != 35 {
		t.Errorf("Expected 35 pages, but got %d", len(chapter.Pages))
	}
}

func TestChapter_Init(t *testing.T) {
	path := "Chapter 0001"
	chapter := Manga.Chapter{Path: path}
	chapter.Init()
}
