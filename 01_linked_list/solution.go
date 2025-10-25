package linkedlist

import (
	"fmt"
	// "os"
	// "reflect"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.count++
	l.tail = &item
}

func (l *LinkedList) Count() int {
	return l.count
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.value == n {
			return *temp, nil
		}
	}

	return Node{value: -1, next: nil}, fmt.Errorf("node with value %d not found", n)
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node

	for temp := l.head; temp != nil; temp = temp.next {
		if temp.value == n {
			nodes = append(nodes, *temp)
		}
	}

	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	l.head = l.DeleteRec(l.head, n, 0, false, all)
}

func (l * LinkedList) DeleteRec(temp *Node, n, delCount int, deleted, all bool) * Node {
	if temp == nil {
		return nil
	}

	if temp.value == n {
		deleted = true
		temp.next = l.DeleteRec(temp.next, n, delCount + 1, deleted, all)
	} else {
		temp.next = l.DeleteRec(temp.next, n, delCount, deleted, all)
	}

	if temp.next == nil {
		l.tail = temp
	}

	if temp.value == n && (all || deleted && delCount == 0) {
		l.count--
		return temp.next
	}

	return temp
}

func (l *LinkedList) Insert(after *Node, add Node) {

}

func (l *LinkedList) InsertFirst(first Node) {

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
	l.count = 0
}
