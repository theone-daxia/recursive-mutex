package recursive_mutex

import (
	"fmt"
	"github.com/petermattis/goid"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的 goroutine id
	recursion int32 // 当前持有锁的 goroutine 的重入次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		// 当前持有锁的 goroutine 就是本次调用的 goroutine，说明是重入
		m.recursion++
		return
	}

	m.Mutex.Lock()
	// 获得锁的 goroutine 第一次调用，记下它的 id，调用次数加一
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		// 非当前持有锁的 goroutine 尝试释放锁，panic
		panic(fmt.Sprintf("curr %d, not owner %d\n", gid, m.owner))
	}

	m.recursion--
	if m.recursion > 0 { // 持有锁的 goroutine 还没完全释放，直接返回
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

/******** 以下为两种获取 goroutine id 的方法 ********/

// 通过 runtime.Stack() 获取 goroutine id
func GoIdByRuntime() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	s := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	idField := strings.Fields(s)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v\n", err))
	}
	return int64(id)
}

/*
通过 hacker 获取 goroutine id

分为以下三步：
	1、获取当前 goroutine 的 TLS 对象
	2、从 TLS 中获取 goroutine 结构的 g 指针
	3、从 g 指针中获取 goroutine id

此处采用第三方库 petermattis/goid
*/
func GoIdByHacker() int64 {
	return goid.Get()
}
