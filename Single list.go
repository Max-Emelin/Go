package main

import (
	"fmt"
	"unsafe"
)

type Ele struct{
	info int
	next *Ele
}
type SingleList struct {
	len int
	head *Ele
}
func (s *SingleList) Print() {
	if s.head==nil{
		fmt.Println("Nothing to print!")
	}else {
		i := 0
		for tmp := s.head; tmp != nil; tmp = tmp.next {
			i++
			fmt.Println(i, "-ый элемент: ", tmp.info)
		}
		fmt.Println()
	}
}
func (s *SingleList) AddFront(info int) {
	tmp := s.head
	s.head = &Ele{info: info}
	s.head.next= tmp
	s.len++
}
func (s *SingleList) AddBack(info int) {
	if s.head==nil{
		s.head=&Ele{info: info,next: nil}

	}else {
		tmp := s.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &Ele{info: info, next: nil}
	}
	s.len++
}
func (s *SingleList) RemoveFront() {
	if s.head!=nil {
		s.head = s.head.next
		s.len--
	}else {
		fmt.Println("Nothing to delete!")
	}
}
func (s *SingleList) RemoveBack() {
	if s.head!=nil {
		tmp:=s.head
		for tmp.next.next!=nil{
			tmp=tmp.next
		}
		tmp.next=nil
		s.len--
	}else {
		fmt.Println("Nothing to delete!")
	}
}
func (s *SingleList) Reverse(){
	if s.head!=nil && s.head.next!= nil {
		tmp := s.head
		for tmp != nil {
			tmp = tmp.next
		}
		curr := s.head
		prev := tmp
		next := s.head.next
		for next != nil {
			curr.next = prev
			prev = curr
			curr = next
			next = curr.next
		}
		curr.next = prev
		s.head = curr
	}else{
		fmt.Println("Noting to reverse!")
	}
}
func (s *SingleList) Size() {
	x:=int(unsafe.Sizeof(s))
	fmt.Println("Size: ",x*s.len)
}
func main() {
	zxc := new(SingleList)
	zxc.Print()
	zxc.RemoveBack()
	zxc.RemoveFront()

	zxc.AddFront(10)
	zxc.Print()

	zxc.AddFront(74)
	zxc.Print()

	zxc.AddBack(4)
	zxc.Print()

	zxc.AddBack(7)
	zxc.Print()

	zxc.Reverse()
	zxc.Size()
	zxc.Print()

	zxc.RemoveBack()
	zxc.RemoveFront()
	zxc.Print()

}
