package ksnowfake

import "github.com/bwmarrin/snowflake"

// NextId get a snowflake id.
func NextId(id *int64,nodeId int64)error {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return  err
	}
	*id=int64(node.Generate())
	return nil
}

// NextIds get snowflake ids.
func NextIds(nodeId,num int64) ([]int64,error) {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return  nil,err
	}
	IdsTmp:= make([]int64, num)
	for i := int64(0); i < num; i++ {
		IdsTmp[i]=int64(node.Generate())
	}
	return IdsTmp, nil
}