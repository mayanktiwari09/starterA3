package distkvs

import (
	"errors"

	"github.com/DistributedClocks/tracing"
)


type ServerStart struct {
	serverId string
}

type ServerJoining struct {
	serverId string
}

type NextServerJoining struct {
	nextServerId string
}

type NewJoinedSuccessor struct {
	nextServerId string
}


type ServerJoined struct {
	serverId string
}


type ServerFailRecvd struct {
	failedServerId string
}


type NewFailoverSuccessor struct {
	newNextServerId string
}

type NewFailoverPredecessor struct {
	newPrevServerId string
}

type ServerFailHandled struct {
	failedServerId string
}


type PutRecvd struct {
	ClientId string
	OpId     uint32
	Key      string
	Value    string
}

type PutOrdered struct {
	ClientId string
	OpId     uint32
	gId      uint32
	Key      string
	Value    string
}


type PutFwd struct {
	ClientId string
	OpId     uint32
	gId      uint32
	Key      string
	Value    string
}

type PutFwdRecvd struct {
	ClientId string
	OpId     uint32
	gId      uint32
	Key      string
	Value    string
}

type PutResult struct {
	ClientId string
	OpId     uint32
	gId      uint32
	Key      string
	Value    string
}


type GetRecvd struct {
	ClientId string
	OpId     uint32
	Key      string
}


type GetOrdered struct {
	ClientId string
	OpId     uint32
	gId      uint32
	Key      string
}

type GetResult struct {
	ClientId string
	OpId     uint32
	Key      string
	Value *string
}

type Server struct {
	// state may go here
}

func (*Server) Start(serverId uint8, coordAddr string, serverAddr string, serverListenAddr string, strace *tracing.Tracer) error {
	return errors.New("not implemented")
}
