package manga

import (
	Misc "main/Misc"
	"os"
	"path/filepath"
)

/*
Representation of a chapter.

# Fields
  - Path:   The path to the chapter folder.
  - Title:  The title of the chapter.
  - Pages:  The list of pages in the chapter.
*/
type Chapter struct {
	Path  string
	Title string
	Pages []Page
}

// ------------------------------
// Getters and Setters
// ------------------------------

/*
Get the path to the chapter folder.

# Returns
  - The path to the chapter folder.
*/
func (c *Chapter) GetPath() string {
	return c.Path
}

/*
Get the title of the chapter.

# Returns
  - The title of the chapter.
*/
func (c *Chapter) GetTitle() string {
	return c.Title
}

/*
Get the list of pages in the chapter.

# Returns
  - The list of pages in the chapter.
*/
func (c *Chapter) GetPages() *[]Page {
	return &c.Pages
}

/*
Set the path to the chapter folder.

# Params
  - path: The path to the chapter folder.
*/
func (c *Chapter) SetPath(path string) {
	c.Path = path
}

/*
Set the title of the chapter.

# Params
  - title: The title of the chapter.

# Note
  - The title will be formatted using Misc.FormatTitle.
*/
func (c *Chapter) SetTitle(title string) {
	c.Title = Misc.FormatTitle(title)
}

/*
Set the list of pages in the chapter.

# Params
  - pages: The list of pages in the chapter.
*/
func (c *Chapter) SetPages(pages *[]Page) {
	c.Pages = *pages
}

// ------------------------------
// Methods
// ------------------------------

/*
Load all pages in the chapter.

# Params
  - basePath: The base path to the manga folder.
*/
func (c *Chapter) LoadPages(basePath string) {
	var pages []Page
	fullPath := filepath.Join(basePath, c.GetPath())

	files, _ := os.ReadDir(fullPath)
	for _, file := range files {
		if !file.IsDir() {
			page := Page{Path: file.Name()}
			page.LoadImageDimensions(fullPath)
			pages = append(pages, page)
		}
	}

	c.SetPages(&pages)
}

/*
Init the chapter.
*/
func (c *Chapter) Init() {
	c.SetTitle(filepath.Base(c.GetPath()))
}
