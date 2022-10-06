package main

import (
	"github.com/ferryvg/go-feeds-aggregator/internal"
)

func main() {
	app := internal.NewApp()

	if err := app.Boot(); err != nil {
		panic(err)
	}

	app.Shutdown()
}
