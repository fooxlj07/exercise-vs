package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	ehlo = "EHLO"
	quit = "QUIT"
	date = "DATE"
)

type action func(msg string)

// ConnectionHandler handle the command and connections
type ConnectionHandler struct {
	conn           net.Conn // connections
	actionsHistory []string // the command history
}

// NewConnectionHandler constructor to create ConnectionHandler instance
func NewConnectionHandler(conn net.Conn) ConnectionHandler {
	conn.Write([]byte("220 localhost \n"))
	log.Println("connect to", conn.RemoteAddr().String())
	return ConnectionHandler{
		conn:           conn,
		actionsHistory: []string{},
	}
}

func (handler *ConnectionHandler) quit(msg string) {
	handler.conn.Write([]byte("221 Bye \n"))
	handler.actionsHistory = append(handler.actionsHistory, quit)
	log.Println("connection will be closed", handler.actionsHistory)
	handler.conn.Close()
}

func (handler *ConnectionHandler) ehlo(msg string) {
	handler.conn.Write([]byte("250 Pleased to meet you " + msg + " \n"))
	handler.actionsHistory = append(handler.actionsHistory, ehlo)
}

func (handler *ConnectionHandler) date(msg string) {
	if len(handler.actionsHistory) > 0 && handler.actionsHistory[len(handler.actionsHistory)-1] == ehlo {
		t := time.Now()
		response := fmt.Sprintf("%d/%d/%dT%d:%d:%d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
		handler.conn.Write([]byte(response + " \n"))
	} else {
		handler.conn.Write([]byte("550 Bad state \n"))
	}

	handler.actionsHistory = append(handler.actionsHistory, date)
}

func (handler *ConnectionHandler) router(command, msg string) {
	switch command {
	case ehlo:
		handler.ehlo(msg)
	case date:
		handler.date(msg)
	case quit:
		handler.quit(msg)
	default:
		handler.conn.Write([]byte("404 Bad Request \n"))
	}
}

func (handler *ConnectionHandler) start() {
	for {
		bufferBytes, err := bufio.NewReader(handler.conn).ReadBytes('\n')
		if err != nil {
			log.Println("client left..")
			handler.conn.Close()
			// escape recursion
			return
		}
		handler.handleMassage(bufferBytes)
	}
}

func (handler *ConnectionHandler) handleMassage(rawMsg []byte) {
	content := strings.Split(strings.TrimSpace(string(rawMsg)), " ")
	command, msg := "", ""
	switch len(content) {
	case 1:
		command = content[0]
	case 2:
		command, msg = content[0], content[1]
	default:
		handler.conn.Write([]byte("404 Bad Request"))
		return
	}
	handler.router(strings.Replace(command, "\n", "", 1), msg)
}
