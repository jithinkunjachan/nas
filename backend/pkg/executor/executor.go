package executor

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/jithinkunjachan/nasserver/backend/pkg/ws"
)

type Builder interface {
	Command() Command
}
type Command struct {
	Cmd  string
	Args []string
	Ws   *ws.WS
}

func Exec(b Builder) error {
	e := b.Command()
	cmd := exec.Command(e.Cmd, e.Args...)
	cmd.Stdin = strings.NewReader("jesus@1save")

	if e.Ws != nil {
		e.Ws.BroadcastJSON(ws.Clear, "")
	}

	cmd.Stderr = cmd.Stdout

	done := make(chan struct{})

	r, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(r)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if e.Ws != nil {
				e.Ws.BroadcastJSON(ws.Message, line)
			}
		}
		done <- struct{}{}
	}()
	err := cmd.Start()
	if err != nil {
		return err
	}

	<-done

	return cmd.Wait()

}
