package command

import (
	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

type Blkid struct {
	Ws *ws.WS
}

func (me *Blkid) Command() executor.Command {
	return executor.Command{
		Cmd: "blkid",
		Ws:  me.Ws,
	}
}

var _ executor.Builder = &Blkid{}
