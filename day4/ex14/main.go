package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
	"sync/atomic"
)

var lock sync.Mutex  // 互斥锁，读写都是互斥的
var rwlock sync.RWMutex  // 读写锁，写是互斥的，读是同时的

func testMutex() {
	var a map[int]int
	a = make(map[int]int, 8)

	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	a[5] = 10
	a[6] = 10
	a[7] = 10

	for i:=0; i< 5; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[2] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

func testRWMutex() {
	var a map[int]int
	a = make(map[int]int, 8)
	var count int32

	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	a[5] = 10
	a[6] = 10
	a[7] = 10

	for i:=0; i< 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[2] = rand.Intn(100)
			time.Sleep(time.Millisecond * 10)
			rwlock.Unlock()
		}(a)
	}

	for i:=0; i< 100; i++ {
		go func(b map[int]int) {
			for {
				rwlock.RLock()
				time.Sleep(time.Millisecond)
				rwlock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(time.Second*2)
	c := atomic.LoadInt32(&count)
	fmt.Println(c)
}

func main() {
	testMutex()
	testRWMutex()
}
