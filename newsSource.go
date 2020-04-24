package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	newsapi "github.com/kaelanb/newsapi-go"
)

type newsClient struct {
	*newsapi.Client
	config *Config
}

func newNewsClient(config *Config) *newsClient {
	client := newsapi.New(config.NewsAPIToken)
	return &newsClient{
		Client: client,
		config: config,
	}
}

func (newsClient *newsClient) checkNews(lastCheck time.Time) []article {
	query1 := []string{"country=" + newsClient.config.Country, "sortBy=publishedAt"}

	newsResponse, err := newsClient.GetTopHeadlines(query1)
	if err != nil {
		log.Println(err)
		return nil
	}

	var newNews []article

	for _, a := range newsResponse.Articles {
		if a.PublishedAt.Unix() > lastCheck.Unix() {
			log.Println("New post:", a.Title)
			newNews = append(newNews, article{
				Article: a,
			})
		} else if newsClient.config.Debug {
			fmt.Println("Old Post:", "["+a.PublishedAt.String()+"]", a.Title)
		}
	}

	sort.Sort(SortByPublished(newNews))

	return newNews
}
