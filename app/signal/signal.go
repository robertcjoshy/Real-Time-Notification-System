package signal

import "errors"

var (
	err = errors.New("no data")
)

type Signal interface {
	Subscribe(id int) (<-chan struct{}, func(), error)
	Publish(id int) error
}
