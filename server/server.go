package server

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type state struct {
	IP         net.Addr
	CurrentDir string
}

// HandleConnection handles incoming connections
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

	case ListDir:
		fNames, err := listdir(connState)
		if err != nil {
			io.WriteString(c, err.Error())
		} else {
			for _, file := range fNames {
				io.WriteString(c, file)
				io.WriteString(c, "\n")
			}
		}

	default:
		io.WriteString(c, CommandNotImplementedError)
	}
}

func listdir(connState *state) ([]string, error) {
	files, err := ioutil.ReadDir(connState.CurrentDir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var fNames []string
	for _, file := range files {
		fNames = append(fNames, file.Name())
	}

	return fNames, nil
}
