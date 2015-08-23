package main

import (
	"log"
	"time"

	"github.com/mem/ws-rpc/internal/common"
)

type Comm struct{}

func (c *Comm) Hello(args *common.HelloArgs, reply *common.HelloReply) error {
	*reply = "Hello!"
	log.Println(args, *reply)
	time.Sleep(1 * time.Second)
	return nil
}
