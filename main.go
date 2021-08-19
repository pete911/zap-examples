package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"go.uber.org/zap/zapcore"
	"sync"
)

func main() {
	logger, _ := NewZapConfig(zapcore.DebugLevel).Build()
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
			a.GetUser(context.WithValue(ctx, "request", generateRequestId()))
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.GetUser(context.WithValue(ctx, "request", generateRequestId()))
		}()
	}
	wg.Wait()
}

// --- helper functions ---

func generateRequestId() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%X-%X", b[0:4], b[4:])
}
