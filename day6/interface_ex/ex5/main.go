package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}

	p.tail.next = node
	p.tail = node
}

func (p *Link) Trans() {
	pp := p.head
	for pp != nil {
		fmt.Println(pp.data)
		pp = pp.next
	}
}

func main() {
	var initLink Link
	for i := 0; i <= 10; i++ {
		initLink.InsertTail(i)
		//initLink.InsertHead(i)
	}

	initLink.Trans()
}
