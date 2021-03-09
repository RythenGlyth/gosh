package debug

import (
	"gosh/src/shared"
	"log"
	"os"
	"strconv"
	"time"

	ipc "github.com/scrouthtv/golang-ipc"
)

// cBufSize is the size of the client's sending buffer.
const cBufSize = 8

// Client can send debugging data from the gosh client to a gosh-debug server.
type Client struct {
	cc        *ipc.Client
	buf       chan *dMsg
	isReading bool
}

type dMsg struct {
	key int
	msg string
}

// StartDebugServer starts a server that gosh instances can connect to
// to send debugging messages.
func StartDebugServer() {
	sc, err := ipc.StartServer("goshdebug", nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("debugger ready, now open gosh with debugging enabled")

	run := true
	for run {
		msg, err := sc.Read()
		if err != nil {
			log.Println(err)
		} else if msg.MsgType > 0 {
			var s string = string(msg.Data)
			if s == "stop" {
				log.Println("going to stop")
				run = false
			} else {
				name := shared.ModuleIdentifierFromInt(msg.MsgType).String()
				log.Printf("%s: %s\n", name, string(msg.Data))
			}
		}
	}
}

// NewClient creates a new client and attaches it to a local gosh-debug server.
func NewClient() (*Client, error) {
	cc, err := ipc.StartClient("goshdebug", nil)
	if err != nil {
		return nil, &LaunchError{err}
	}

	c := &Client{cc, make(chan *dMsg, cBufSize), false}

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

	c.cc.Write(shared.ModMain.AsInt(), []byte("Connected from pid "+strconv.Itoa(os.Getpid())))

	var msg *dMsg

	for {
		msg = <-c.buf
		c.cc.Write(msg.key, []byte(msg.msg))
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
func (c *Client) SendMessage(k shared.ModuleIdentifier, msg string) {
	if c == nil || c.cc == nil {
		return
	}
	c.buf <- &dMsg{k.AsInt(), msg}
}
