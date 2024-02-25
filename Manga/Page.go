package manga

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

/*
Representation of a page.

# Fields
  - Path:  The path to the page file.
  - W:	 The width of the image.
  - H:	 The height of the image.
*/
type Page struct {
	Path string
	W    int
	H    int
}

// ------------------------------
// Getters and Setters
// ------------------------------

/*
Get the path to the page file.

# Returns
  - The path to the page file.
*/
func (p *Page) GetPath() string {
	return p.Path
}

/*
Get the width of the image.

# Returns
  - The width of the image.
*/
func (p *Page) GetW() int {
	return p.W
}

/*
Get the height of the image.

# Returns
  - The height of the image.
*/
func (p *Page) GetH() int {
	return p.H
}

/*
Set the path to the page file.

# Params
  - path: The path to the page file.
*/
func (p *Page) SetPath(path string) {
	p.Path = path
}

/*
Set the width of the image.

# Params
  - w: The width of the image.
*/
func (p *Page) SetW(w int) {
	p.W = w
}

/*
Set the height of the image.

# Params
- h: The height of the image.
*/
func (p *Page) SetH(h int) {
	p.H = h
}

// ------------------------------
// Methods
// ------------------------------

/*
Load image dimensions.

# Params
  - basePath: The base path to the page file.
*/
func (p *Page) LoadImageDimensions(basePath string) {
	file, err := os.Open(filepath.Join(basePath, p.GetPath()))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		panic(err)
	}

	p.SetW(img.Width)
	p.SetH(img.Height)
}
