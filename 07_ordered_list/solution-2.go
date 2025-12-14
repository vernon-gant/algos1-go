package ordered_list

func (l *OrderedList[T]) RemoveDuplicates() {
	if l.head == nil || l.head == l.tail {
		return
	}

	for temp := l.head; temp != nil; {
		innerTemp := temp.next
		for ; innerTemp != nil && innerTemp.value == temp.value; innerTemp = innerTemp.next {
			l.count--
		}
		temp.next = innerTemp
		if innerTemp != nil {
			innerTemp.prev = temp
		}
		temp = innerTemp
	}
}