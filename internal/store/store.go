package store

import (
	"context"
	"errors"
	"time"
)

var ErrConflict = errors.New("data conflict")

// MessageStore описывает абстрактное хранилище сообщений пользователей
type MessageStore interface {
	FindRecipient(ctx context.Context, username string) (userID string, err error)
	ListMessages(ctx context.Context, userID string) ([]Message, error)
	GetMessage(ctx context.Context, id int64) (*Message, error)
	SaveMessage(ctx context.Context, userID string, msg Message) error
	RegisterUser(ctx context.Context, userID, username string) error
}

// Message описывает объект сообщения
type Message struct {
	ID      int64     // внутренний идентификатор сообщения
	Sender  string    // отправитель
	Time    time.Time // время отправления
	Payload string    // текст сообщения
}
