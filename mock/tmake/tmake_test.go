package main

import "fmt"

type Oo struct {
	K []int64

}

func main()  {
	tmp:=Oo{K: make([]int64,0)}
	if tmp.K==nil {
		fmt.Println("哇塞。。。。")
	}
}
