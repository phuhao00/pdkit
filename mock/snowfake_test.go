package mock

import (
	"pdkit/tea/klogger"
	"pdkit/tea/krpc"
	"pdkit/tea/ksignal"
	"pdkit/tea/ksnowfake"
	"testing"
)

func TestSnowFake(t *testing.T) {
	rcvr:=&ksnowfake.SnowflakeRPC{}
	logger:=klogger.DefaultLogger
	err:=krpc.InitRPC(rcvr,"127.0.0.1:8090",logger)
	if err != nil {
		logger.ErrorF("bind rpc err ")
	}
	sg:=ksignal.InitSignal()
	hsg:=ksignal.HandleSignal(sg,nil)
	hsg()
}
