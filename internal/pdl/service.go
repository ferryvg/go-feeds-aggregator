package pdl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
)

type Service struct {
	store          *Store
	loader         Loader
	updateTicker   *time.Ticker
	stopCh         chan struct{}
	updateInterval time.Duration
}

func NewService() *Service {
	return &Service{
		store:  NewStore(),
		loader: NewLoader(),
	}
}

func (s *Service) Init() error {
	if err := s.loadFeed(); err != nil {
		return err
	}

	go s.watchUpdate()

	return nil
}

func (s *Service) Stop() {
	s.stopCh <- struct{}{}
}

func (s *Service) Load(ctx context.Context, limit int) ([]*domain.FeedItem, error) {
	if limit <= 0 {
		return nil, nil
	}

	result := make([]*domain.FeedItem, 0, limit)

	items := s.store.GetMulti(limit)
	for _, item := range items {
		feedItem, err := s.convertToDomain(item)
		if err != nil {
			return nil, err
		}

		result = append(result, feedItem)
	}

	return result, nil
}

func (s *Service) loadFeed() error {
	feed, err := s.loader.Load(context.Background())
	if err != nil {
		return err
	}

	s.refreshUpdateInterval(feed)
	s.updateStore(feed)

	return nil
}

func (s *Service) refreshUpdateInterval(feed *RssFeedXml) {
	updateInterval := time.Hour

	switch strings.TrimSpace(feed.Channel.UpdatePeriod) {
	case "daily":
		updateInterval = time.Hour * 24
	case "weekly":
		updateInterval = time.Hour * 24 * 7
	case "monthly":
		updateInterval = time.Hour * 24 * 30
	case "yearly":
		updateInterval = time.Hour * 24 * 365
	}

	s.updateInterval = updateInterval * time.Duration(feed.Channel.UpdateFrequency)

	if s.updateTicker != nil {
		s.updateTicker.Reset(s.updateInterval)
	}
}

func (s *Service) watchUpdate() {
	ticker := time.NewTicker(s.updateInterval)

	for {
		select {
		case <-s.stopCh:
			return

		case <-ticker.C:
			if err := s.loadFeed(); err != nil {
				fmt.Printf("[ERROR] Failed to refresh feed: %s\n", err.Error())
			}
		}
	}
}

func (s *Service) updateStore(feed *RssFeedXml) {
	s.store.Update(feed.Channel.Items)
}

func (s *Service) convertToDomain(item *RssItem) (*domain.FeedItem, error) {
	pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pub date '%s': %w", item.PubDate, err)
	}

	var content string
	if item.Content != nil {
		content = item.Content.Content
	}

	return &domain.FeedItem{
		Title:       item.Title,
		Description: content,
		ContentUrl:  item.Link,
		PublishDate: pubDate,
	}, nil
}
