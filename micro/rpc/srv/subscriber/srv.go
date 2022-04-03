package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	srv "day1/micro/rpc/srv/proto/srv"
)

type Srv struct{}

func (e *Srv) Handle(ctx context.Context, msg *srv.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *srv.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
