package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/robert/notification/app/entity"
	"github.com/robert/notification/app/signal"
	"github.com/robert/notification/app/storage"
)

type Bobol struct {
	storage       storage.Storage
	signal        signal.Signal
	defualtimeout time.Duration
}

func Newbolbol(str storage.Storage, sig signal.Signal) *Bobol {
	return &Bobol{
		storage:       str,
		signal:        sig,
		defualtimeout: 2 * time.Minute,
	}
}

var (
	mut            sync.Mutex
	hasuserfetched = make(map[int]bool, 0)
	lastseen       = make(map[int]time.Time)
)

func Build() *Bobol {
	fmt.Println("inside build function")
	str := storage.Newmemorywithlist(100)
	sig := signal.Newtopic()
	return Newbolbol(str, sig)
}

/*
	func (b *Bobol) Getnotifications(ctx context.Context, clientid int, timestamp int64) ([]entity.Notification, error) {
		c, err := b.storage.Count(ctx, clientid)
		if err != nil {
			return nil, fmt.Errorf("no value in count")
		}
		if c > 0 {
			return b.storage.Popall(ctx, clientid)
		}
		ch, close, err := b.signal.Subscribe(clientid)
		defer close()
		if err != nil {
			return nil, fmt.Errorf("cannnot retrive channel in subcribe")
		}
		ctx, ctxcancel := context.WithTimeout(ctx, b.defualtimeout)
		defer ctxcancel()

		select {
		case <-ch:
			return b.storage.Pop(ctx, clientid, timestamp)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
*/
func (b *Bobol) Getnotifications(ctx context.Context, clientID int, timestamp int64) ([]entity.Notification, error) {

	mut.Lock()
	isFirstRequest := !hasuserfetched[clientID]
	if isFirstRequest {
		hasuserfetched[clientID] = true
	}
	mut.Unlock()

	if isFirstRequest {
		notifications, err := b.storage.Popall(ctx, clientID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		lastseen[clientID] = time.Now()
		fmt.Println("popall = ", notifications)
		return notifications, nil
	}

	// If this is not the first request, subscribe to the signal channel for new notifications
	ch, close, err := b.signal.Subscribe(clientID)
	defer close()
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to channel: %v", err)
	}

	// Wait for either a notification or a timeout
	ctx, cancel := context.WithTimeout(ctx, b.defualtimeout)
	defer cancel()

	select {
	case <-ch:
		// If a notification is received, fetch and return new notifications
		return b.storage.Pop(ctx, clientID, &mut, lastseen)
	case <-ctx.Done():
		// If timeout occurs, return error indicating timeout
		return nil, ctx.Err()
	}
}

func (b *Bobol) Notify(ctx context.Context, clientid int, notification entity.Notification) error {
	fmt.Println("inside function Notify")
	err := b.storage.Push(ctx, clientid, notification)
	fmt.Println("after push operation")
	if err != nil {
		return fmt.Errorf("cannot push")
	}
	errr := b.signal.Publish(clientid)
	if errr != nil {
		fmt.Println("ERROR IN notify = ", errr)
		return errr
	}
	return nil
}
