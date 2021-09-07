package main

import (
	"context"
	"github.com/pete911/zap-examples/logger"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

func main() {

	logger, err := logger.NewZapConfig(zapcore.DebugLevel).Build()
	if err != nil {
		log.Fatalf("build zap logger: %v", err)
	}

	store := NewStore(logger)

	serviceA := NewServiceA(logger, store)
	serviceB := NewServiceB(logger, store)
	run(context.Background(), serviceA, serviceB)
}

func run(ctx context.Context, a ServiceA, b ServiceB) {

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.GetUser(ctx)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.GetUser(ctx)
		}()
	}
	wg.Wait()
}

