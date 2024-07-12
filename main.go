//go:build !test
// +build !test

package main

import (
	"webmalc/shutterstock-downloader/cmd"
	"webmalc/shutterstock-downloader/common/config"
	downloader "webmalc/shutterstock-downloader/internal"

	"webmalc/shutterstock-downloader/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	runner := downloader.NewDownloader(log)
	config.Setup()
	router := cmd.NewCommandRouter(runner)
	router.Run()
}
