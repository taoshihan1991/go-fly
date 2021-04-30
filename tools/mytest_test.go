package tools

import (
	"sync"
	"testing"
)

func TestMyTest(t *testing.T) {
	MyTest()
}
func TestMyStruct(t *testing.T) {
	MyStruct()
}

type SMap struct {
	sync.RWMutex
	Map map[int]int
}

func (l *SMap) readMap(key int) (int, bool) {
	l.RLock()
	value, ok := l.Map[key]
	l.RUnlock()
	return value, ok
}

func (l *SMap) writeMap(key int, value int) {
	l.Lock()
	l.Map[key] = value
	l.Unlock()
}

var mMap *SMap

func TestMyMap(t *testing.T) {
	mMap = &SMap{
		Map: make(map[int]int),
	}

	for i := 0; i < 10000; i++ {
		go func() {
			mMap.writeMap(i, i)
		}()
		go readMap(i)
	}
}
func readMap(i int) (int, bool) {
	return mMap.readMap(i)
}
