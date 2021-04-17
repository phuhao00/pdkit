package mock

import (
	"fmt"
	"net/rpc"
	"pdkit/tea/ksnowfake"
	"testing"
	"time"
)

func TestRpcClient(t *testing.T) {
	clt, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf(err.Error())
	}
	//tmpStop := make(chan bool, 1)
	//pingAndRetry(tmpStop, clt, "127.0.0.1:8090")
	Ids(2,clt)

}

const (
	zkNodeDelaySleep    = 1 * time.Second // zk error delay sleep
	zkNodeDelayChild    = 3 * time.Second // zk node delay get children
	rpcClientPingSleep  = 1 * time.Second // rpc client ping need sleep
	rpcClientRetrySleep = 1 * time.Second // rpc client retry connect need sleep

	RPCPing    = "SnowflakeRPC.Ping"
	RPCNextId  = "SnowflakeRPC.NextId"
	RPCNextIds = "SnowflakeRPC.NextIds"
)

// pingAndRetry ping the rpc connect and re connect when has an error.
func  pingAndRetry(stop <-chan bool, client *rpc.Client, addr string) {
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("client.Close() error(%v)", err)
		}
	}()
	var (
		failed bool
		status int
		err    error
		tmp    *rpc.Client
	)
	for {
		select {
		case <-stop:
			fmt.Printf("addr: \"%s\" pingAndRetry goroutine exit", addr)
			return
		default:
		}
		if !failed {
			if err = client.Call(RPCPing, 0, &status); err != nil {
				fmt.Printf("client.Call(%s) error(%v)", RPCPing, err)
				failed = true
				continue
			} else {
				failed = false
				time.Sleep(rpcClientPingSleep)
				continue
			}
		}
		if tmp, err = rpc.Dial("tcp", addr); err != nil {
			fmt.Printf("rpc.Dial(tcp, %s) error(%v)", addr, err)
			time.Sleep(rpcClientRetrySleep)
			continue
		}
		client = tmp
		failed = false
		fmt.Printf("client reconnect %s ok", addr)
	}
}

//
func Ids(nodeId int64,client *rpc.Client) (ids []int64, err error) {
	args:=&ksnowfake.Args{
		NodeId: nodeId,
		Num: 5,
	}
	ids=make([]int64,0)
	if err = client.Call(RPCNextIds, args, &ids); err != nil {
		fmt.Printf("rpc.Call(\"%s\", &id) error(%v)", RPCNextId, err)
	}
	fmt.Println(args, ids)
	return
}