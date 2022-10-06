package xkcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
	"github.com/valyala/fasthttp"
)

const (
	sourceUrlPattern = "https://xkcd.com/%s/info.0.json"
	requestTimeout   = time.Second * 5
)

type Service struct {
	cache  Cache
	client *fasthttp.Client
}

func NewService() *Service {
	return &Service{
		cache:  NewCache(),
		client: new(fasthttp.Client),
	}
}

func (s *Service) Init() error {
	return nil
}

func (s *Service) Stop() {
}

func (s *Service) Load(ctx context.Context, limit int) ([]*domain.FeedItem, error) {
	if limit <= 0 {
		return nil, nil
	}

	sourceItems := make([]*SourceItem, 0, limit)

	var nextId int

iter:
	for i := 0; i < limit; i++ {
		var (
			item *SourceItem
			err  error
		)

		switch i {
		case 0:
			item, err = s.loadRemoteById(0)

		default:
			item, err = s.getById(nextId)
		}

		if err != nil {
			return nil, err
		}

		if item == nil {
			continue iter
		}

		sourceItems = append(sourceItems, item)
		s.cache.Set(item)

		nextId = item.Id - 1
	}

	return s.convertToDomain(sourceItems), nil
}

func (s *Service) loadRemoteById(id int) (*SourceItem, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	url := s.buildRemoteUrl(id)
	req.SetRequestURI(url)

	if err := s.client.DoTimeout(req, resp, requestTimeout); err != nil {
		return nil, err
	}

	if resp.StatusCode() == fasthttp.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(fmt.Sprintf(
			"unexpected remote response status code: %d; details: %s",
			resp.StatusCode(), string(resp.Body()),
		))
	}

	sourceItem := new(SourceItem)
	if err := json.Unmarshal(resp.Body(), sourceItem); err != nil {
		return nil, fmt.Errorf("unexpected or invalid remote response: %w", err)
	}

	return sourceItem, nil
}

func (s *Service) buildRemoteUrl(id int) string {
	var idStr string
	if id > 0 {
		idStr = strconv.FormatInt(int64(id), 10)
	}

	return fmt.Sprintf(sourceUrlPattern, idStr)
}

func (s *Service) convertToDomain(sourceItems []*SourceItem) []*domain.FeedItem {
	result := make([]*domain.FeedItem, 0, len(sourceItems))

	for _, item := range sourceItems {
		result = append(result, item.ToFeed())
	}

	return result
}

func (s *Service) getById(id int) (*SourceItem, error) {
	item, err := s.cache.Get(id)
	if err == nil {
		return item, nil
	}

	if !errors.Is(NotFoundErr, err) {
		return nil, err
	}

	return s.loadRemoteById(id)
}
