package krpc

import (
	"net"
	"net/rpc"
	"pdkit/tea/klogger"
)

// rpcListen start rpc listen.
func rpcListen(bind string,logger *klogger.KLogger) {
	l, err := net.Listen("tcp", bind)
	if err != nil {
		logger.ErrorF("net.Listen(\"tcp\", \"%s\") error(%v)", bind, err)
		panic(err)
	}
	// if process exit, then close the rpc bind
	defer func() {
		logger.InfoF("rpc addr: \"%s\" close", bind)
		if err := l.Close(); err != nil {
			logger.ErrorF("listener.Close() error(%v)", err)
		}
	}()
	rpc.Accept(l)
}

// InitRPC  start rpc listen.
func InitRPC(rcvr interface{},bind string,logger *klogger.KLogger) error {
	err:=rpc.Register(rcvr)
	if err != nil {
		return err
	}
	logger.InfoF("start listen rpc addr: \"%s\"", bind)
	go rpcListen(bind,logger)
	return nil
}
