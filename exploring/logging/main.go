package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "https://google.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "https://google.com")

	logText := `
	This is a text that will be logged
	with a lot of information
	and a lot of details
	and a lot of context
	and a lot of metadata
	and a lot of tags
	and a lot of labels
	and a lot of fields
	and a lot of values
	`

	sugar.Infow("logText", "text", logText)
}
