package rpcstore

import (
	"context"
	"github.com/ipfs/go-cid"
	"sync"
)

var taskPool = &sync.Pool{
	New: func() any {
		return new(task)
	},
}

func getTask() *task {
	return taskPool.Get().(*task)
}

type result struct {
	cid cid.Cid
	err error
}

type task struct {
	ctx         context.Context
	peerId      string
	spaceId     string
	cid         cid.Cid
	denyPeerIds []string
	write       bool
	exec        func(c *client) error
	onFinished  func(t *task, c *client, err error)
	ready       chan result
}

func (t *task) execWithClient(c *client) {
	t.onFinished(t, c, t.exec(c))
}

func (t *task) release() {
	t.ctx = nil
	t.peerId = ""
	t.spaceId = ""
	t.denyPeerIds = t.denyPeerIds[:0]
	t.write = false
	t.exec = nil
	taskPool.Put(t)
}
