package app

import (
	"context"
	"fmt"
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

func Build() *Bobol {
	fmt.Println("inside build function")
	str := storage.Newmemorywithlist(100)
	sig := signal.Newtopic()
	return Newbolbol(str, sig)
}

func (b *Bobol) Getnotifications(ctx context.Context, clientid int) ([]entity.Notification, error) {
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
		return b.storage.Popall(ctx, clientid)
	case <-ctx.Done():
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
		return errr
	}
	return nil
}
