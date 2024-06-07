package storage

import (
	"context"
	"fmt"
	"sync"
	"time"

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
	fmt.Println("inside push")
	item := m.get(clientid)
	fmt.Println("AFTER get")
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

func (m *Memorywithlist) Pop(ctx context.Context, clientid int, mut *sync.Mutex, lastseen map[int]time.Time) ([]entity.Notification, error) {
	item := m.get(clientid)
	var notification []entity.Notification
	if len(item.notifications) == 0 {
		return nil, Errempty
	}

	mut.Lock()
	tiime, ok := lastseen[clientid]
	if !ok {
		tiime = time.Time{}
	}
	mut.Unlock()
	item.mu.Lock()
	for _, value := range item.notifications {
		nottime := value.(entity.Messagenotification).Createdat
		data := time.Unix(nottime, 0)
		if data.After(tiime) {
			notification = append(notification, value.(entity.Messagenotification))
			//return nil, Errempty
		}
	}
	item.mu.Unlock()

	//notification := item.notifications[0]

	//item.notifications = item.notifications[1:]
	fmt.Println("pop = ", notification)
	return notification, nil
}

func (m *Memorywithlist) Popall(ctx context.Context, clientid int) ([]entity.Notification, error) {
	item := m.get(clientid)

	if len(item.notifications) == 0 {
		return nil, Errempty
	}

	item.mu.Lock()
	defer item.mu.Unlock()
	/*
		defer func() {
			item.notifications = nil
		}()
	*/
	return item.notifications, nil

}

func (m *Memorywithlist) get(clientid int) *Userstorage {
	item, _ := m.storage.LoadOrStore(clientid, &Userstorage{mu: new(sync.Mutex)})
	return item.(*Userstorage)
}
