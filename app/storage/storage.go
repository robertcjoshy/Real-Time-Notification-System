package storage

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/robert/notification/app/entity"
)

var Errempty = errors.New("'no notification")

type Storage interface {
	Push(ctx context.Context, clientid int, notification entity.Notification) error
	Count(ctx context.Context, clientid int) (int, error)
	Pop(ctx context.Context, clientid int, mut *sync.Mutex, lastseen map[int]time.Time) ([]entity.Notification, error)
	Popall(ctx context.Context, clientid int) ([]entity.Notification, error)
}
