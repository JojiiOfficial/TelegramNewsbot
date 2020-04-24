package main

import (
	"fmt"
	"strings"

	"github.com/JesusIslam/tldr"
	newsapi "github.com/kaelanb/newsapi-go"
)

type article struct {
	newsapi.Article
}

// SortByPublished sorts article by publish date
type SortByPublished []article

func (a SortByPublished) Len() int      { return len(a) }
func (a SortByPublished) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByPublished) Less(i, j int) bool {
	return a[i].PublishedAt.Unix() < a[j].PublishedAt.Unix()
}

func (article *article) getShortDescription() string {
	bag := tldr.New()
	result, err := bag.Summarize(article.Description, 1)
	if len(result) == 0 || err != nil {
		if err != nil {
			fmt.Println(err)
		}
		return article.Description
	}

	return result[0]
}

func (article *article) getTitle() string {
	if strings.Contains(article.Title, "-") {
		return strings.TrimSpace(article.Title[:strings.LastIndex(article.Title, "-")])
	}

	return strings.TrimSpace(article.Title)
}

func (article *article) String() string {
	return fmt.Sprintf("%s\n\n%s\n\n%s", article.getTitle(), article.getShortDescription(), article.URL)
}
