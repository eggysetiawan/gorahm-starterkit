package main

import (
	"github.com/eggysetiawan/gorahm-starterkit/config"
	"github.com/eggysetiawan/gorahm-starterkit/app"
)

func main() {
	config.LoadEnv(".env")

	app.Router()

	// app.Console()
}
