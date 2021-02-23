package debug

import (
	"fmt"
	"os"
	"strconv"
	"time"

	ipc "github.com/scrouthtv/golang-ipc"
)

// StartDebugServer starts a server that gosh instances can connect
// to to send debugging messages
func StartDebugServer() {
	sc, err := ipc.StartServer("goshdebug", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	run := true
	for run {
		msg, err := sc.Read()
		if err != nil {
			fmt.Println(err)
		} else if msg.MsgType > 0 {
			var s string = string(msg.Data)
			if s == "stop" {
				fmt.Println("going to stop")
				run = false
			} else {
				fmt.Printf("%d: %s\n", msg.MsgType, string(msg.Data))
			}
		}
	}
}

// Client can be used to send debugging data from the gosh client to a gosh-debug server
type Client struct {
	cc        *ipc.Client
	buf       chan *ipc.Message
	isReading bool
}

// NewClient creates a new client and attaches it to a local gosh-debug server.
func NewClient() (*Client, error) {
	cc, err := ipc.StartClient("goshdebug", nil)
	if err != nil {
		return nil, err
	}

	var c *Client = &Client{cc, make(chan *ipc.Message, 8), false}

	go c.readLoop()
	go c.writeLoop()

	return c, nil
}

func (c *Client) status() ipc.Status {
	var sl []ipc.Status = []ipc.Status{
		ipc.NotConnected, ipc.Listening, ipc.Connecting, ipc.Connected,
		ipc.ReConnecting, ipc.Closed, ipc.Closing, ipc.Error, ipc.Timeout,
	}

	var ms string = c.cc.Status()
	for _, s := range sl {
		if ms == s.String() {
			return s
		}
	}
	return -1
}

func (c *Client) writeLoop() {
	var status ipc.Status = c.status()
	for status != ipc.Connected {
		if status == ipc.Closed || status == ipc.Closing || status == ipc.Error || status == ipc.Timeout {
			os.Stdout.WriteString("\r\n\r\nCould not connect to ipc connect\r\n\r\n")
			os.Exit(1)
		}
		time.Sleep(100 * time.Millisecond)
		status = c.status()
	}
	c.cc.Write(1, []byte("Connected from pid "+strconv.Itoa(os.Getpid())))
	var msg *ipc.Message
	for {
		msg = <-c.buf
		c.cc.Write(msg.MsgType, msg.Data)
	}
}

func (c *Client) readLoop() {
	for {
		_, err := c.cc.Read()
		if err != nil {
			return
		}
	}
}

// SendMessage writes a debugging message to the attached debugger.
// If this client isn't connected, the function returns.
func (c *Client) SendMessage(k int, msg string) {
	if c == nil || c.cc == nil {
		return
	}
	c.buf <- &ipc.Message{MsgType: k, Data: []byte(msg)}
}
