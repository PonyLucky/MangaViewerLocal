package manga_test

import (
	Manga "main/Manga"
	"testing"
)

func TestManga_GetPath(t *testing.T) {
	path := "The Other World’s Wizard Does Not Chant"
	manga := Manga.Manga{Path: path}
	expectedPath := path

	actualPath := manga.GetPath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestManga_GetTitle(t *testing.T) {
	manga := Manga.Manga{Title: "The Other World’s Wizard Does Not Chant"}
	expectedTitle := "The Other World’s Wizard Does Not Chant"

	actualTitle := manga.GetTitle()

	if *actualTitle != expectedTitle {
		t.Errorf("Expected title: %s, but got: %s", expectedTitle, *actualTitle)
	}
}

func TestManga_GetChapters(t *testing.T) {
	chapter1 := Manga.Chapter{Path: "Chapter 0001"}
	chapter2 := Manga.Chapter{Path: "Chapter 0002"}
	manga := Manga.Manga{Chapters: []Manga.Chapter{chapter1, chapter2}}
	expectedChapters := []Manga.Chapter{chapter1, chapter2}

	actualChapters := manga.GetChapters()

	if len(*actualChapters) != len(expectedChapters) {
		t.Errorf("Expected %d chapters, but got %d", len(expectedChapters), len(*actualChapters))
	}

	for i, chapter := range expectedChapters {
		if (*actualChapters)[i].Path != chapter.Path {
			t.Errorf("Expected chapter at index %d: %s, but got: %s", i, chapter.Path, (*actualChapters)[i].Path)
		}
	}
}

func TestManga_GetPoster(t *testing.T) {
	manga := Manga.Manga{Poster: "poster.jpg"}
	expectedPoster := "poster.jpg"

	actualPoster := manga.GetPoster()

	if string(*actualPoster) != expectedPoster {
		t.Errorf("Expected poster: %s, but got: %s", expectedPoster, *actualPoster)
	}
}

func TestManga_SetPath(t *testing.T) {
	manga := Manga.Manga{}
	expectedPath := "The Other World’s Wizard Does Not Chant"

	manga.SetPath(expectedPath)

	if manga.Path != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, manga.Path)
	}
}

func TestManga_SetTitle(t *testing.T) {
	manga := Manga.Manga{}
	expectedTitle := "The Other World’s Wizard Does Not Chant"

	manga.SetTitle(expectedTitle)

	if manga.Title != expectedTitle {
		t.Errorf("Expected title: %s, but got: %s", expectedTitle, manga.Title)
	}
}

func TestManga_SetChapters(t *testing.T) {
	chapter1 := Manga.Chapter{Path: "Chapter 0001"}
	chapter2 := Manga.Chapter{Path: "Chapter 0002"}
	manga := Manga.Manga{}
	expectedChapters := []Manga.Chapter{chapter1, chapter2}

	manga.SetChapters(&expectedChapters)

	if len(manga.Chapters) != len(expectedChapters) {
		t.Errorf("Expected %d chapters, but got %d", len(expectedChapters), len(manga.Chapters))
	}

	for i, chapter := range expectedChapters {
		if manga.Chapters[i].Path != chapter.Path {
			t.Errorf("Expected chapter at index %d: %s, but got: %s", i, chapter.Path, manga.Chapters[i].Path)
		}
	}
}

func TestManga_SetPoster(t *testing.T) {
	manga := Manga.Manga{}
	expectedPoster := "poster.jpg"

	manga.SetPoster(&expectedPoster)

	if manga.Poster != expectedPoster {
		t.Errorf("Expected poster: %s, but got: %s", expectedPoster, manga.Poster)
	}
}

func TestManga_LoadChapters(t *testing.T) {
	manga := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	expectedChapters := []Manga.Chapter{
		{Path: "Chapter 0001"},
		{Path: "Chapter 0002"},
	}

	manga.LoadChapters("/home/louis/Documents/Hakuneko/Mangas/")

	actualChapters := manga.GetChapters()
	var tmpChapters []Manga.Chapter

	if len(*actualChapters) > 2 {
		tmpChapters = (*actualChapters)[:2]
	}

	if len(tmpChapters) != len(expectedChapters) {
		t.Errorf("Expected %d chapters, but got %d", len(expectedChapters), len(tmpChapters))
	}

	for i, chapter := range expectedChapters {
		if (tmpChapters)[i].Path != chapter.Path {
			t.Errorf("Expected chapter at index %d: %s, but got: %s", i, chapter.Path, (tmpChapters)[i].Path)
		}
	}
}

func TestManga_LoadPoster(t *testing.T) {
	manga := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	expectedPoster := "poster.jpg"

	manga.LoadPoster("/home/louis/Documents/Hakuneko/Mangas/")

	actualPoster := manga.GetPoster()

	if string(*actualPoster) != expectedPoster {
		t.Errorf("Expected poster: %s, but got: %s", expectedPoster, *actualPoster)
	}
}

func TestManga_GetChapter(t *testing.T) {
	chapter1 := Manga.Chapter{Path: "Chapter 0001"}
	chapter2 := Manga.Chapter{Path: "Chapter 0002"}
	manga := Manga.Manga{Chapters: []Manga.Chapter{chapter1, chapter2}}
	expectedChapter := chapter1

	actualChapter := manga.GetChapter(0)

	if actualChapter.Path != expectedChapter.Path {
		t.Errorf("Expected chapter: %s, but got: %s", expectedChapter.Path, actualChapter.Path)
	}
}

func TestManga_GetNumberOfChapters(t *testing.T) {
	chapter1 := Manga.Chapter{Path: "Chapter 0001"}
	chapter2 := Manga.Chapter{Path: "Chapter 0002"}
	manga := Manga.Manga{Chapters: []Manga.Chapter{chapter1, chapter2}}
	expectedNumberOfChapters := 2

	actualNumberOfChapters := manga.GetNumberOfChapters()

	if actualNumberOfChapters != expectedNumberOfChapters {
		t.Errorf("Expected number of chapters: %d, but got: %d", expectedNumberOfChapters, actualNumberOfChapters)
	}
}

func TestManga_Init(t *testing.T) {
	manga := Manga.Manga{Path: "The Other World’s Wizard Does Not Chant"}
	manga.Init("/home/louis/Documents/Hakuneko/Mangas/")
}
