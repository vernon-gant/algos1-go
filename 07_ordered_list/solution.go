package ordered_list

import (
	"constraints"
	"errors"

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
	} else if l.shouldBeInserted(l.head, item, true) {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else if l.shouldBeInserted(l.tail, item, false) {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	} else {
		toInsert := l.head
		for ; l.shouldBeInserted(toInsert, item, true); toInsert = toInsert.next {	}
		newNode.next = toInsert
		newNode.prev = toInsert.prev
		toInsert.prev.next = newNode
		toInsert.prev = newNode
	}
	l.count++
}

func (l *OrderedList[T]) shouldBeInserted(node *Node[T], item T, before bool) bool {
	if before {
		return (l._ascending && item <= node.value) || (!l._ascending && item >= node.value)
	}
	return (l._ascending && item >= node.value) || (!l._ascending && item <= node.value)
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	var result Node[T]

	if l.head == nil {
		return result, errors.New("empty list")
	}

	if n < l.head.value || n > l.tail.value {
		return result, errors.New("not found")
	}

	
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
