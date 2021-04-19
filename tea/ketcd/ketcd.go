package ketcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type KEtcd struct {
	*clientv3.Client
	requestTimeout time.Duration
	eps []string
}

var DefaultDialTimeOut = 3*time.Second

//NewKEtcd create KEtcd
func NewKEtcd(endpoints []string,dialTimeout time.Duration) *KEtcd {
	return &KEtcd{
		Client:         GetEtcdClient(endpoints,dialTimeout),
		requestTimeout: DefaultDialTimeOut,
	}
}

//Put put key value normal
func (k *KEtcd)Put(key,value string)error {
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	_, err:= k.Client.Put(ctx, key, value)
	cancel()
	return err
}

//Get get normal
func (k *KEtcd)Get(key string) (*clientv3.GetResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	resp, err := k.Client.Get(ctx, key)
	cancel()
	return resp, err
}

//GetWithRev get with revision
func  (k *KEtcd)GetWithRev(key string,revision int64)(*clientv3.GetResponse,error) {
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	resp, err := k.Client.Get(ctx, key, clientv3.WithRev(revision))
	cancel()
	return resp, err
}

//GetSortedPrefix get sorted with prefix
func (k *KEtcd)GetSortedPrefix(key string)(*clientv3.GetResponse,error) {
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	resp, err := k.Client.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	return resp, err
}

// Del  delete key with prefix
func  (k *KEtcd)Del(key string) (string,error) {
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	defer cancel()
	// count keys about to be deleted
	gresp, err := k.Client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return "", err
	}
	// delete the keys
	dresp, err :=  k.Client.Delete(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return "", err
	}
	return  fmt.Sprint("Deleted all keys:", int64(len(gresp.Kvs)) == dresp.Deleted),err
}

//KVCompact kv specify compact revision of your choice
func  (k *KEtcd)KVCompact(key string)(*clientv3.CompactResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), k.requestTimeout)
	resp, err := k.Client.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}
	compRev := resp.Header.Revision // specify compact revision of your choice

	ctx, cancel = context.WithTimeout(context.Background(), k.requestTimeout)
	cpResp, err:=  k.Client.Compact(ctx, compRev)
	cancel()
	return cpResp,err
}

//KVTxn
/*
func (kv clientv3.KV)error{
_, err = kvc.Put(context.TODO(), "key", "xyz")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = kvc.Txn(ctx).
		// txn value comparisons are lexical
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")).
		// the "Then" runs, since "xyz" > "abc"
		Then(clientv3.OpPut("key", "XYZ")).
		// the "Else" does not run
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()
	if err != nil {
		return err
	}
}
*/
func (k *KEtcd)KVTxn(f func(kv clientv3.KV)error) error{
	kvc := clientv3.NewKV(k.Client)
	return f(kvc)
}

//KVDo
//Do applies a single Op on KV without a transaction.
//Do is useful when creating arbitrary operations to be issued at a
//later time; the user can range over the operations, calling Do to
//execute them. Get/Put/Delete, on the other hand, are best suited
//for when the operation should be issued at the time of declaration.
func  (k *KEtcd)KVDo(ctx context.Context,ops []clientv3.Op) error {
	for _, op := range ops {
		if _, err := k.Client.Do(ctx, op); err != nil {
			return err
		}
	}
	return nil
}


