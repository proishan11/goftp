package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type state struct {
	IP         net.Addr
	CurrentDir string
}

// HandleConnection hancles incomming connections
func HandleConnection(c net.Conn) {

	defer c.Close()
	currdir, err := os.Getwd()
	io.WriteString(c, ConnectionSuccessMessage)

	if err != nil {
		log.Fatal(err)
		io.WriteString(c, OSError)
		return
	}
	// Add an authentication layer here

	input := bufio.NewScanner(c)
	connState := &state{
		IP:         c.RemoteAddr(),
		CurrentDir: currdir,
	}

	for input.Scan() {
		cmd := formatString(input.Text())
		handleCommand(cmd, connState, c)
	}

}

func formatString(s string) string {
	return strings.TrimSpace(s)
}

func handleCommand(cmd string, connState *state, c net.Conn) {
	switch cmd {
	case PresentDir:
		io.WriteString(c, connState.CurrentDir)
		io.WriteString(c, "\n")
	default:
		io.WriteString(c, CommandNotImplementedError)
	}
}
