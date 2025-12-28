package dictionary

import (
	"constraints"
	"errors"
	"github.com/vernon-gant/algos1-go/07_ordered_list"
)

/*
* 9. Dictionary - task number 5 - dictionary using Ordered List
*
* I decided to take existing implementation of the OrderedList(more precisely its interface aka spec) and use it as internal field
* and not to reimplement the list from scratch. We take the interface, although in the case of go the concrete struct from previous
* lesson - but this does not matter, because concrete implementations may differ in efficincy. Our previous implementation was O(n)
* for search and insertion, but using an array for this would result in O(log(n)) for search. Specification could also (probably) mention
* that these operations must result in O(log(n)). This could make sense.
* In our case we typically want from the spec level that an Ordered List has search in O(n) or even O(log(n)) and maybe also insertion.
* Both log(n) would be an implementation using a skip list, the second one would be just dynamic array. But the interface is still the same!
* So we just take the OrderedList and use its operations. For that I added another method FindPosition which could be implemented
* in any data structure aka skip list, doubly linked list, dyn array. It returns bool and int where bool indicates that element was found
* or not and int its position which the element needs to be inserted to or is at. Because we need to shift all the value elements, the Put
* complexity is O(n) - copy does not do that in O(1). Same applies to Delete. Seach, however, depends on the implementation of the
* FindPosition.
*/

type OrderedDict[K constraints.Ordered, V any] struct {
	keys   *ordered_list.OrderedList[K]
	values []V
}

func NewOrderedDict[K constraints.Ordered, V any]() *OrderedDict[K, V] {
	keys := &ordered_list.OrderedList[K]{}
	keys.Clear(true)
	return &OrderedDict[K, V]{
		keys:   keys,
		values: make([]V, 0),
	}
}

func (d *OrderedDict[K, V]) Put(key K, value V) {
	position, found := d.keys.FindPosition(key)

	if found {
		d.values[position] = value
		return
	}

	d.keys.Add(key)

	var zero V

	d.values = append(d.values, zero)
	copy(d.values[position+1:], d.values[position:])
	d.values[position] = value
}

func (d *OrderedDict[K, V]) Delete(key K) error {
	position, found := d.keys.FindPosition(key)

	if !found {
		return errors.New("key not found")
	}

	d.keys.Delete(key)

	copy(d.values[position:], d.values[position+1:])
	d.values = d.values[:len(d.values)-1]

	return nil
}

func (d *OrderedDict[K, V]) IsKey(key K) bool {
	_, found := d.keys.FindPosition(key)
	return found
}