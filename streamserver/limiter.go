package main

import "log"

/**
流量控制
bucket: token 1
token 2 ... token n
request -> 发送一个token
然后返回response, 再返回一个token
类似于一种信号量机制的?
 */

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool{
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <- cl.bucket
	log.Println("New Connection coming: ", c)
}