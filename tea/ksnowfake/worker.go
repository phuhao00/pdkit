package ksnowfake

import (
	"pdkit/tea/klogger"
	"time"
)

type SnowflakeRPC struct {
	logger klogger.KLogger
}

type Args struct {
	Num int64
	NodeId int64
}

func (r *SnowflakeRPC) NextId(Id int64, nodeId *int64) error {
	return NextId(&Id, *nodeId)
}

func (r *SnowflakeRPC) NextIds( args *Args,ids *[]int64) error {
	tmpIds,err:=NextIds(args.NodeId, args.Num)
	if err!=nil {
		return err
	}
	*ids=tmpIds
	return nil
}

// Ping return the service status.
func (r *SnowflakeRPC) Ping(ignore int, status *int) error {
	*status = 0
	return nil
}

// Timestamp return the service current unix seconds.
func (s *SnowflakeRPC) Timestamp(ignore int, timestamp *int64) error {
	*timestamp = time.Now().Unix()
	return nil
}
