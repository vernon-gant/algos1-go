package main

import (
	"strconv"
	// "os"
)

type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	return 0
}

func (ht *HashTable) SeekSlot(value string) int {
	return -1
}

func (ht *HashTable) Put(value string) int {
	return -1
}

func (ht *HashTable) Find(value string) int {
	return -1
}
