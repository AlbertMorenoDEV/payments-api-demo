package main

import (
	"context"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/infrastructure"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/ui"
	"log"
	"net/http"
	"sync"
)

const serverPort = 1234

func main() {
	domainEvents := make(chan domain.DomainEvent)
	config := infrastructure.NewConfig()
	services := infrastructure.BuildServices(config, domainEvents)
	router := ui.BuildRouter(services)

	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go startWebServer(wg, router)
	startConsumers(rootCtx, wg, services)

	wg.Wait()
}

func startWebServer(wg *sync.WaitGroup, router http.Handler) {
	defer wg.Done()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
}

func startConsumers(ctx context.Context, wg *sync.WaitGroup, services *infrastructure.Services) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		services.BalanceProjector().Start(ctx)
	}()
}
