package revercity

import (
	"fmt"
	"sync"
)

type Control struct {
	toStop      bool
	totalBytes  int
	fails       int
	connections int
	mut         sync.Mutex
}

func (c *Control) Stop() {
	c.toStop = true
}

func (c *Control) Connections() int {
	return c.connections
}

func (c *Control) TotalBytes() int {
	return c.totalBytes
}

func (c *Control) Fails() int {
	return c.fails
}

func newControl() *Control {
	return &Control{
		toStop:     false,
		totalBytes: 0,
		fails:      0,
		mut:        sync.Mutex{},
	}
}

func (c *Control) String() string {
	var stopStatus = "Running"
	if c.toStop {
		stopStatus = "Stopped"
	}
	return fmt.Sprintf("Control[%s -> Connections: %d  Bytes: %d  Fails: %d]",
		stopStatus,
		c.Connections(),
		c.TotalBytes(),
		c.Fails(),
	)
}
