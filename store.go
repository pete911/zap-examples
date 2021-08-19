package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

var usernames = []string{"peter", "bob", "john", "joseph", "daniel"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Store struct {
	logger *zap.Logger
}

func NewStore(logger *zap.Logger) Store {
	return Store{logger: logger.With(zap.String("component", "store"))}
}

func (s Store) GetUser(ctx context.Context) User {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	s.logger.Debug("get user", GetContext(ctx)...)
	user := User{
		Id:       generateId(),
		Username: generateUsername(),
		Password: "super-secret",
	}
	s.logger.Debug(fmt.Sprintf("user %s found", user), GetContext(ctx)...)
	return user
}

// --- helper functions ---

func generateId() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%X", b)
}

func generateUsername() string {
	return usernames[rand.Intn(len(usernames))]
}
