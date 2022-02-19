package distkvs


import (
	"errors"
	"github.com/DistributedClocks/tracing"
)

// Actions to be recorded by coord (as part of ctrace, ktrace, and strace):

type CoordStart struct {
	
}

type ServerFail struct {
	serverId string
}

type Put struct {
	ClientId string
	OpId     uint32
	Key      string
	Value    string
}

type Get struct {
	ClientId string
	OpId     uint32
	Key      string
}

type ServerFailHandledRecvd struct {
	failedServerId uint32
	adjacentServerId uint32
	
}

type NewChain struct {
	chain  []uint32
}

type HeadReqRecvd struct {
	ClientId string
}

type HeadRes struct {
	ClientId string
	ServerId string
}

type TailReqRecvd struct {
	ClientId string
	ServerId string
}

type TailRes struct {
	ClientId string
	ServerId string
}

type ServerJoiningRecvd struct {
	ServerId string
}

type ServerJoinedRecvd struct {
	ServerId string
}


type AllServersJoined struct {

}


// // NotifyChannel is used for notifying the client about a mining result.
// type NotifyChannel chan ResultStruct

// type ResultStruct struct {
// 	opId        uint32
// 	StorageFail bool
// 	Result      *string
// }

// type KVS struct {
// 	notifyCh NotifyChannel
// 	// Add more KVS instance state here.
// }

// func NewKVS() *KVS {
// 	return &KVS{
// 		notifyCh: nil,
// 	}
// }


type Coord struct {
	// state may go here
}

func (*Coord) Start(clientAPIListenAddr string, serverAPIListenAddr string, lost_msgs_thresh uint8, numServers uint8, ctrace *tracing.Tracer) error {
	return errors.New("not implemented")
}

