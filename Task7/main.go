package main

import (
	"fmt"
	"sync"
	"time"
)

type SomeMap struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewMap() *SomeMap {
	return &SomeMap{
		m: make(map[int]int),
	}
}

func (s *SomeMap) Write(key int, value int) {
	s.mx.Lock()
	s.m[key] = value
	s.mx.Unlock()
}

func (s *SomeMap) Read(key int) (int, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func f1(key, val int, m *SomeMap) {

	for i, j := key, val; i < key+100; i++ {
		m.Write(i, j)
		j++
	}
}

func f2(key int, m *SomeMap) {

	for i := key; i < key+100; i++ {
		fmt.Println(m.Read(i))
	}
}

func main() {
	map1 := NewMap()
	go f1(0, 0, map1)
	go f2(0, map1)
	time.Sleep(1 * time.Second)
}
