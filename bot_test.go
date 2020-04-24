package main

import (
	"testing"

	"github.com/kaelanb/newsapi-go"
)

func TestArticleTitle(t *testing.T) {
	rawTitle := "Der Börsen-Tag:14:34 BMW fährt schrittweise Produktion hoch"
	article := article{newsapi.Article{Title: rawTitle + " - n-tv NACHRICHTEN"}}

	formattedTitle := article.getTitle()

	if formattedTitle != rawTitle {
		t.Errorf("Error formatting title: got '%s' but expected '%s'", formattedTitle, rawTitle)
	}
}
