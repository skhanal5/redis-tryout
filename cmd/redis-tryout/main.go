package main

import (
	"flag"

	"github.com/alecthomas/kong"
	"github.com/skhanal5/redis-tryout/internal"
	"github.com/skhanal5/redis-tryout/internal/cache"
	"github.com/skhanal5/redis-tryout/internal/command"
)

func main() {
	config := &internal.Options{}
	flag.StringVar(&config.RedisAddr, "RedisHost", "localhost:6379", "Redis server host")
	client := cache.NewCache(*config)
	cliCtx := command.AppContext{
		Cache: &client,
	}
	// Bind the AppContext to the kong parser
	ctx := kong.Parse(&command.Commands{})
	ctx.Bind(cliCtx)
	err := ctx.Run(cliCtx)
	ctx.FatalIfErrorf(err)
}
