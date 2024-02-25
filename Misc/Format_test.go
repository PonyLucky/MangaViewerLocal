package misc_test

import (
	Misc "main/Misc"
	"testing"
)

func TestFormatTitle(t *testing.T) {
	title := "the other world’s wizard does not chant"
	expectedTitle := "The Other World’s Wizard Does Not Chant"

	actualTitle := Misc.FormatTitle(title)

	if actualTitle != expectedTitle {
		t.Errorf("Expected title: %s, but got: %s", expectedTitle, actualTitle)
	}
}
