package ordered_list

import (
	"constraints"
	// "os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	count      int
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	return l.count
}

func (l *OrderedList[T]) Add(item T) {
	newNode := &Node[T]{value: item}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else if (l._ascending && item <= l.head.value) || (!l._ascending && item >= l.head.value) {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else if (l._ascending && item >= l.tail.value) || (!l._ascending && item <= l.tail.value) {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	} else {
		toInsert := l.head
		for ; (l._ascending && item >= toInsert.value) || (!l._ascending && item <= toInsert.value); toInsert = toInsert.next {}
		newNode.next = toInsert
		newNode.prev = toInsert.prev
		toInsert.prev.next = newNode
		toInsert.prev = newNode
	}
	l.count++
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	return Node[T]{value: n, next: nil, prev: nil}, nil
}

func (l *OrderedList[T]) Delete(n T) {

}

func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l.count = 0
	l._ascending = asc
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}
