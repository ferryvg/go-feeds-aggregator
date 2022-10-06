package domain

import "time"

type FeedItem struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PictureUrl  string    `json:"pictureUrl"`
	ContentUrl  string    `json:"contentUrl"`
	PublishDate time.Time `json:"publishDate"`
}
