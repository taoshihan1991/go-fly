package tools

import (
	"log"
	"sync"
	"time"
)

type LimitQueeMap struct {
	sync.RWMutex
	LimitQueue map[string][]int64
}

func (l *LimitQueeMap) readMap(key string) ([]int64, bool) {
	l.RLock()
	value, ok := l.LimitQueue[key]
	l.RUnlock()
	return value, ok
}

func (l *LimitQueeMap) writeMap(key string, value []int64) {
	l.Lock()
	l.LimitQueue[key] = value
	l.Unlock()
}

var LimitQueue = &LimitQueeMap{
	LimitQueue: make(map[string][]int64),
}
var ok bool

func NewLimitQueue() {
	cleanLimitQueue()
}
func cleanLimitQueue() {
	go func() {
		for {
			log.Println("cleanLimitQueue start...")
			LimitQueue.LimitQueue = nil
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//单机时间滑动窗口限流法
func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	if LimitQueue.LimitQueue == nil {
		LimitQueue.LimitQueue = make(map[string][]int64)
	}
	if _, ok = LimitQueue.readMap(queueName); !ok {
		LimitQueue.writeMap(queueName, make([]int64, 0))
		return true
	}
	q, _ := LimitQueue.readMap(queueName)
	//队列未满
	if uint(len(q)) < count {
		LimitQueue.writeMap(queueName, append(q, currTime))
		return true
	}
	//队列满了,取出最早访问的时间
	earlyTime := q[0]
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		q = q[1:]
		LimitQueue.writeMap(queueName, append(q, currTime))
	}
	return true
}
