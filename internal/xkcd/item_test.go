package xkcd_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
	"github.com/ferryvg/go-feeds-aggregator/internal/xkcd"
)

var (
	PublishYear  = 2000
	PublishMonth = 01
	PublishDay   = 01

	CorrectFeedItem = &domain.FeedItem{
		Title:       "Woodpecker",
		Description: "[[A man with a beret and a woman are standing on a boardwalk]]",
		PictureUrl:  "https://imgs.xkcd.com/comics/woodpecker.png",
		ContentUrl:  "",
		PublishDate: time.Date(PublishYear, time.Month(PublishMonth), PublishDay, 0, 0, 0, 0, time.UTC),
	}
)

func TestSourceItem_ToFeed(t *testing.T) {
	tests := []struct {
		name string
		want *domain.FeedItem
	}{
		{
			name: "success",
			want: CorrectFeedItem,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := xkcd.SourceItem{
				SafeTitle:   CorrectFeedItem.Title,
				Description: CorrectFeedItem.Description,
				ImageUrl:    CorrectFeedItem.PictureUrl,
				ContentUrl:  CorrectFeedItem.ContentUrl,
				Year:        PublishYear,
				Month:       PublishMonth,
				Day:         PublishDay,
			}

			result := item.ToFeed()
			assert.EqualValues(t, tt.want, result)
		})
	}
}
