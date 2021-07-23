package main

import (
	"context"

	"github.com/Toolnado/quote-service.git/internal/onederx"
)

func main() {
	source := onederx.NewSource()

	source.Start(context.Background())
}
