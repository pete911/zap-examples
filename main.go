package main

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

func main() {

	logger, err := NewZapConfig(zapcore.DebugLevel).Build()
	if err != nil {
		log.Fatalf("build zap logger: %v", err)
	}

	store := NewStore(logger)

	serviceA := NewServiceA(logger, store)
	serviceB := NewServiceB(logger, store)
	run(context.Background(), logger, serviceA, serviceB)
}

func run(ctx context.Context, logger *zap.Logger, a ServiceA, b ServiceB) {

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := a.GetUser(ctx); err != nil {
				logger.Error("get user", zap.Error(err))
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := b.GetUser(ctx); err != nil {
				logger.Error("get user", zap.Error(err))
			}
		}()
	}
	wg.Wait()
}
