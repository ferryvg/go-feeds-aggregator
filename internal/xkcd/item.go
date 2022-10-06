package xkcd

import (
	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
	"time"
)

type SourceItem struct {
	Id          int    `json:"num"`
	ContentUrl  string `json:"link"`
	Day         int    `json:"day,string"`
	Month       int    `json:"month,string"`
	Year        int    `json:"year,string"`
	Title       string `json:"title"`
	SafeTitle   string `json:"safe_title"`
	ImageUrl    string `json:"img"`
	Description string `json:"alt"`
}

func (i *SourceItem) ToFeed() *domain.FeedItem {
	date := time.Date(i.Year, time.Month(i.Month), i.Day, 0, 0, 0, 0, time.UTC)

	return &domain.FeedItem{
		Title:       i.SafeTitle,
		Description: i.Description,
		PictureUrl:  i.ImageUrl,
		ContentUrl:  i.ContentUrl,
		PublishDate: date,
	}
}
