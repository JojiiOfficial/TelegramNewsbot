package main

import (
	"testing"

	"github.com/kaelanb/newsapi-go"
)

func TestArticleTitle(t *testing.T) {
	rawTitle := "I'am a nice and interesting title"
	article := article{newsapi.Article{Title: rawTitle + "- jojii"}}

	formattedTitle := article.getTitle()

	if formattedTitle != rawTitle {
		t.Errorf("Error formatting title: got '%s' but expected '%s'", formattedTitle, rawTitle)
	}
}
