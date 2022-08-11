package main

import (
	"context"

	"github.com/ansedo/gophermart/internal/logger"
	"github.com/ansedo/gophermart/internal/server"
	"github.com/ansedo/gophermart/internal/services/shutdowner"
)

func main() {
	logger.New()
	server.Run(context.Background())
	<-shutdowner.Get().ChShutdowned
}
