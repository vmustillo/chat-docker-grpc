package main

import (
	"os"

	"chat-docker-grpc/proto"

	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

// Connection - Data structure used for sending message streams
type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

// Server - Holds a slice of Connections
type Server struct {
	Connection []*Connection
}

// CreateStream - Used by Server struct to create streams
// Creates a connection, appends the connection to the servers connetion slice, and returns a channel error
func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func main() {

}
