package pdl

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	sourceUrl      = "https://feeds.feedburner.com/PoorlyDrawnLines"
	requestTimeout = time.Second * 5
)

type Loader interface {
	Load(ctx context.Context) (*RssFeedXml, error)
}

type loaderImpl struct {
	client         *fasthttp.Client
	stringsCleaner *strings.Replacer
}

func NewLoader() Loader {
	return &loaderImpl{
		client:         new(fasthttp.Client),
		stringsCleaner: strings.NewReplacer("\n", "", "\t", ""),
	}
}

func (l *loaderImpl) Load(ctx context.Context) (*RssFeedXml, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(sourceUrl)

	if err := l.client.DoTimeout(req, resp, requestTimeout); err != nil {
		return nil, fmt.Errorf("failed to load rss feed: %w", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(fmt.Sprintf(
			"unexpected remote response status code: %d; details: %s",
			resp.StatusCode(), string(resp.Body()),
		))
	}

	feed := new(RssFeedXml)
	if err := xml.Unmarshal(resp.Body(), feed); err != nil {
		return nil, fmt.Errorf("unexpected or invalid remote response: %w", err)
	}

	feed.Channel.UpdatePeriod = strings.TrimSpace(l.stringsCleaner.Replace(feed.Channel.UpdatePeriod))

	return feed, nil
}
