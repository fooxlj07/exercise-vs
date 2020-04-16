package main

import (
	"fmt"
	"testing"
	"time"
	"vadesecure/tcp-server/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestConneciton(t *testing.T) {
	conn := new(mocks.ConnectionMock)
	conn.On("Write", mock.Anything).Return(0, nil)

	NewConnectionHandler(conn)
	conn.AssertCalled(t, "Write", []byte("220 localhost \n"))

}

func TestCommand(t *testing.T) {
	conn := new(mocks.ConnectionMock)
	conn.On("Write", mock.Anything).Return(0, nil)
	conn.On("Read", mock.Anything).Return(0, nil)
	conn.On("Close").Return(nil)
	connectionHandler := NewConnectionHandler(conn)
	now := time.Now()
	tt := []struct {
		payload []byte
		want    []byte
	}{
		{
			[]byte("EHLO lol"),
			[]byte("250 Pleased to meet you lol \n"),
		},
		{
			[]byte("DATE"),
			[]byte(fmt.Sprintf("%d/%d/%dT%d:%d:%d \n", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())),
		},
		{
			[]byte("DATE"),
			[]byte("550 Bad state \n"),
		},
		{
			[]byte("whatever"),
			[]byte("404 Bad Request \n"),
		},
		{
			[]byte("QUIT"),
			[]byte("221 Bye \n"),
		},
	}

	for _, tc := range tt {
		connectionHandler.handleMassage(tc.payload)
		conn.AssertCalled(t, "Write", tc.want)
	}
	assert.Equal(t, []string{"EHLO", "DATE", "DATE", "QUIT"},
		connectionHandler.actionsHistory)

	conn.AssertNumberOfCalls(t, "Write", 6)
}
