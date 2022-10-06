package internal

import (
	"encoding/json"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/ferryvg/go-feeds-aggregator/internal/pdl"
	"github.com/ferryvg/go-feeds-aggregator/internal/sources"
	"github.com/ferryvg/go-feeds-aggregator/internal/xkcd"
	"github.com/valyala/fasthttp"
)

var itemsLimit = 20

type App interface {
	Boot() error
	Shutdown()
}

type appImpl struct {
	router      *fasthttprouter.Router
	feedAdapter sources.SourceAdapter
}

func NewApp() App {
	return &appImpl{}
}

func (a *appImpl) Boot() error {
	a.router = fasthttprouter.New()

	if err := a.initAdapter(); err != nil {
		return fmt.Errorf("failed to initialise adapter: %w", err)
	}

	a.router.GET("/", a.feedHandle)

	if err := fasthttp.ListenAndServe(":8080", a.router.Handler); err != nil {
		return fmt.Errorf("failed to initialize server: %w", err)
	}

	return nil
}

func (a *appImpl) Shutdown() {
}

func (a *appImpl) feedHandle(ctx *fasthttp.RequestCtx) {
	items, err := a.feedAdapter.Load(ctx, itemsLimit)
	if err != nil {
		a.internalErr(ctx)
		return
	}

	response, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		a.internalErr(ctx)
		return
	}

	ctx.Response.SetBody(response)
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json")
}

func (a *appImpl) initAdapter() error {
	a.feedAdapter = sources.NewCompoundAdapter(xkcd.NewService(), pdl.NewService())

	return a.feedAdapter.Init()
}

func (a *appImpl) internalErr(ctx *fasthttp.RequestCtx) {
	ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
	ctx.Done()
}
