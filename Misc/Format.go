package misc

import (
	"strings"
)

/*
Format the title of a manga.

# Example

	```
	title := "THE REINCARNATION MAGICIAN OF THE INFERIOR EYES"
	fmt.Println(FormatTitle(title))
	// The Reincarnation Magician Of The Inferior Eyes
	```

# Params
- title: The title of the manga.

# Returns
- The formatted title of the manga.
*/
func FormatTitle(title string) string {
	// Split the title into words.
	newTitle := ""
	isCap := true

	for _, char := range title {
		if char == ' ' {
			newTitle += " "
			isCap = true
		} else {
			if isCap {
				newTitle += strings.ToUpper(string(char))
				isCap = false
			} else {
				newTitle += strings.ToLower(string(char))
			}
		}
	}

	return newTitle
}
