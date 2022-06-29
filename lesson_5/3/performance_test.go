// 3. Протестируйте производительность операций чтения и записи на множестве
// действительных чисел, безопасность которого обеспечивается sync.Mutex и
// sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50%
// запись, 50% чтение; 90% запись, 10% чтение

package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Set struct {
	sync.Map
}

func NewSet() *Set {
	return &Set{
		Map: sync.Map{},
	}
}
func (s *Set) Add(i int) {
	s.Map.Store(i, struct{}{})
}
func (s *Set) Has(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}

func BenchmarkSet(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixMicro())
	var ds = make([]int, 1e8)
	for i := 0; i < 1e8; i += 1 {
		ds[i] = rand.Intn(100)
	}

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(p *testing.PB) {
			i := 0
			for p.Next() {
				if ds[i] < 90 {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i += 1
			}
		})
		b.RunParallel(func(p *testing.PB) {
			i := 0
			for p.Next() {
				if ds[i] < 50 {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i += 1
			}
		})
		b.RunParallel(func(p *testing.PB) {
			i := 0
			for p.Next() {
				if ds[i] < 10 {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i += 1
			}
		})
	})
}
