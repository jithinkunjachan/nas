package command

import (
	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

type Lsblk struct {
	Ws *ws.WS
}

var _ executor.Builder = &Lsblk{}

func (me *Lsblk) Command() executor.Command {
	return executor.Command{
		Cmd:  "lsblk",
		Args: nil,
		Ws:   me.Ws,
	}
}
