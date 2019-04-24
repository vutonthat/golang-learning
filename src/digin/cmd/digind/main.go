package main

import (
	"digin/adapters/ipinfo"
	"digin/server"
	"digin/services"
	"go.uber.org/dig"
	"log"
)

const (
	IPInfoBaseURL = "http://ipinfo.io"
)

func main() {
	err := digContainer().Invoke(func(s *server.Server) {
		defer s.Close()
		if err := s.Run(); err != nil {
			log.Fatalf("Running server: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Revoking server: %v", err)
	}
}

func digContainer() *dig.Container {
	container := dig.New()
	{
		_ = container.Provide(func() ipinfo.Adapter { return ipinfo.NewAdapter(IPInfoBaseURL) })
		_ = container.Provide(services.NewPublicIPService)
		_ = container.Provide(server.New)
	}
	return container
}
