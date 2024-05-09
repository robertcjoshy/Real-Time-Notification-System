package storage

import (
	"context"
	"sync"

	"github.com/robert/notification/app/entity"
)

type Userstorage struct {
	mu            *sync.Mutex
	notifications []entity.Notification
}

type Memorywithlist struct {
	size    int
	storage *sync.Map
}

func Newmemorywithlist(size int) Storage {
	return &Memorywithlist{
		size:    size,
		storage: new(sync.Map),
	}
}

func (m *Memorywithlist) Push(ctx context.Context, clientid int, notification entity.Notification) error {
	item := m.get(clientid)
	item.mu.Lock()
	defer item.mu.Unlock()

	if len(item.notifications) == m.size {
		item.notifications = item.notifications[1:]
	}

	item.notifications = append(item.notifications, notification)
	return nil
}

func (m *Memorywithlist) Count(ctx context.Context, clientid int) (int, error) {
	item := m.get(clientid)
	return len(item.notifications), nil
}

func (m *Memorywithlist) Pop(ctx context.Context, clientid int) (entity.Notification, error) {
	item := m.get(clientid)

	if len(item.notifications) == 0 {
		return nil, Errempty
	}

	item.mu.Lock()
	defer item.mu.Unlock()

	notification := item.notifications[0]
	item.notifications = item.notifications[1:]
	return notification, nil
}

func (m *Memorywithlist) Popall(ctx context.Context, clientid int) ([]entity.Notification, error) {
	item := m.get(clientid)

	if len(item.notifications) == 0 {
		return nil, Errempty
	}

	item.mu.Lock()
	defer item.mu.Unlock()

	defer func() {
		item.notifications = nil
	}()

	return item.notifications, nil

}

func (m *Memorywithlist) get(clientid int) *Userstorage {
	item, _ := m.storage.LoadOrStore(clientid, &Userstorage{mu: new(sync.Mutex)})
	return item.(*Userstorage)
}
