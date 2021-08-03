package stack

import (
	"sync"
)

/*
	Стэк - очередь, которая работает по принципу FIFO(last in frist out - последний вошел послединй вышел).
	Данные, которые были положены последние - берутся первыми, принцип как будто складываешь тарелки,и когда нужно взять,
	ты берешь самую верхнюю.

	Реализует три метода
		Pop - извлечение превого на очередь элемента
		Push - добавление в очередь элемента
		Peek - вывод вершины очереди (последнего элемента)
*/
type Stack struct {
	val   map[int]interface{}
	count int
	mt    sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		val:   map[int]interface{}{},
		count: 0,
	}
}

func (s *Stack) Push(elem interface{}) {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.val[s.Size()] = elem

	s.count++
}

func (s *Stack) Pop() interface{} {
	s.mt.Lock()
	defer s.mt.Unlock()
	if s.Size() == 0 {
		return nil
	}

	val, ok := s.val[s.Size()-1]
	if !ok {
		return nil
	}

	delete(s.val, s.Size()-1)
	s.count--
	return val
}

func (s *Stack) Peek() interface{} {
	s.mt.Lock()
	defer s.mt.Unlock()
	if s.Size() == 0 {
		return nil
	}

	val, ok := s.val[s.Size()-1]
	if !ok {
		return nil
	}

	return val
}

func (s *Stack) Size() int {
	return s.count
}
