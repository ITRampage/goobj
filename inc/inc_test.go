package inc_test

import (
	"sync"
	"testing"

	. "github.com/ITRampage/goobj/inc"
)

// It shows how object/messaging-oriented code performs
func BenchmarkObj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := New(0)
		wg := &sync.WaitGroup{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go (func() {
				defer wg.Done()
				for k := 0; k < 10000; k++ {
					o.Inc()
				}
			})()
		}
		wg.Wait()
	}
}

// It shows a performance of simple incrementation just for
// comparison and understanding overheads of messaging, concurrency and locking
func BenchmarkInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		for ; c < 1000000; c++ {
		}
	}
}

// Because of no locking the code above will generate
// random value for `c` variable at the end of each
// outermost iteration.
//
// This code shows why do we need objects or locking
// and allows to understand the locking overhead.
func BenchmarkCIncNoLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		wg := &sync.WaitGroup{}
		//m := &sync.Mutex{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go (func() {
				defer wg.Done()
				for k := 0; k < 10000; k++ {
					//m.Lock()
					c++
					//m.Unlock()
				}
			})()
		}
		wg.Wait()
	}
}

// Is shows how concurrent incrementation with locking performs.
func BenchmarkCInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		wg := &sync.WaitGroup{}
		m := &sync.Mutex{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go (func() {
				defer wg.Done()
				for k := 0; k < 10000; k++ {
					m.Lock()
					c++
					m.Unlock()
				}
			})()
		}
		wg.Wait()
	}
}

// It shows how concurrent incrementation with locking performs
// when we incrementations in a bulk instead of for every single one incrementation.
//
// The example above allows to better estimate an overhead of locking. Here we make only 100
// locks instead of 1_000_000 locks as in BenchmarkCInc().
func BenchmarkCIncLockAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := 0
		wg := &sync.WaitGroup{}
		m := &sync.Mutex{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go (func() {
				defer wg.Done()
				m.Lock()
				for k := 0; k < 10000; k++ {
					c++
				}
				m.Unlock()
			})()
		}
		wg.Wait()
	}
}
