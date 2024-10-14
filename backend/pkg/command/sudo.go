package command

import (
	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

type Sudo struct {
	Ws   *ws.WS
	Cmd  string
	Args []string
}

func (me *Sudo) Command() executor.Command {
	return executor.Command{
		Cmd:  me.Cmd,
		Args: me.Args,
		Ws:   me.Ws,
	}
}

var _ executor.Builder = &Sudo{}
