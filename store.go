package main

import (
	"context"
	"fmt"
	"github.com/pete911/zap-examples/logger"
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

func (s Store) GetUser(ctx context.Context) (User, error) {

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	s.logger.Debug("get user", logger.GetLoggerFields(ctx)...)
	user := User{
		Id:       generateId(),
		Username: generateUsername(),
		Password: "super-secret",
	}
	s.logger.Debug(fmt.Sprintf("user %s found", user), logger.GetLoggerFields(ctx)...)
	return user, nil
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
