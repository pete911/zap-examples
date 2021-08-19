package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

// --- service A ---

type ServiceA struct {
	logger *zap.Logger
	store  Store
}

func NewServiceA(logger *zap.Logger, store Store) ServiceA {
	return ServiceA{
		logger: logger.With(zap.String("component", "service-a")),
		store:  store,
	}
}

func (s ServiceA) GetUser(ctx context.Context) User {
	s.logger.Debug("get user", GetContext(ctx)...)
	user := s.store.GetUser(ctx)
	s.logger.Debug(fmt.Sprintf("user %s", user), GetContext(ctx)...)
	return user
}

// --- service B ---

type ServiceB struct {
	logger *zap.Logger
	store  Store
}

func NewServiceB(logger *zap.Logger, store Store) ServiceB {
	return ServiceB{
		logger: logger.With(zap.String("component", "service-b")),
		store:  store,
	}
}

func (s ServiceB) GetUser(ctx context.Context) User {
	s.logger.Debug("get user", GetContext(ctx)...)
	user := s.store.GetUser(ctx)
	s.logger.Debug(fmt.Sprintf("user %s", user), GetContext(ctx)...)
	return user
}
