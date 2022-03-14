package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Number int
	FullName string
	Position string
	Next *Node
	Prev *Node
}
type DoubleLinkedList struct {
	len int
	Head *Node
	Tail *Node
}
type Iterator struct {
	Node *Node
}

func (i *Iterator) HasNext() bool {
	if i.Node.Next!=nil{
		return  true
	}else{
		return false
	}
}
func (i *Iterator) Next() (*Iterator,error) {
	if i.HasNext()==true {
		i.Node = i.Node.Next
		return i,nil
	}else{
		return nil,errors.New("No next pls <3")
	}
}

func (i *Iterator) HasPrev() bool {
	if i.Node.Prev!=nil{
		return  true
	}else{
		return false
	}
}
func (i *Iterator) Prev() (*Iterator,error) {
	if i.HasPrev()==true {
		i.Node = i.Node.Prev
		return i,nil
	}else{
		return nil,errors.New("No next pls <3")
	}
}

func BeautyString(begin int,len int, end int) (string) {
	space:=""
	for i:=0;i<(end-begin-len);i++{
		space+=" "
	}
	space+="│"
	return space
}
func BeautyInt(begin int,i int,end int) string {
	lengh:=1
	for i>=10{
		i/=10
		lengh++
	}
	return (BeautyString(begin,lengh,end))
}

func (d *DoubleLinkedList) Print()  {
	if d.Head.Next==d.Tail{
		fmt.Println("Nothing to print!")
	}else{
		fmt.Println("\n│NodeNumber\t  │Number\t   │FullName\t\t\t│Position\t\t\t│")
		i:=0
		RoflNN:=""
		RoflN:=""
		RoflFN:=""
		RoflP:=""
		for tmp:=d.Head.Next;tmp!=d.Tail;tmp=tmp.Next{
			i++
			RoflNN=BeautyInt(3,i,13)
			RoflN=BeautyInt(17,tmp.Number,27)
			RoflFN=BeautyString(27,len(tmp.FullName),45)
			RoflP=BeautyString(45,len(tmp.Position),62)
			fmt.Println("│ ",i,RoflNN,
				tmp.Number,RoflN,
				tmp.FullName,RoflFN,
				tmp.Position,RoflP)
		}
		fmt.Println()
	}
}
func (d *DoubleLinkedList) PushFront(Number int,FullName string,Position string){
	tmp:=d.Head.Next
	d.Head.Next=&Node{Number: Number,FullName: FullName,Position: Position}
	d.Head.Next.Next=tmp
	d.Head.Next.Prev=d.Head
	d.Head.Next.Next.Prev=d.Head.Next
	d.len++
}
func (d *DoubleLinkedList) PushBack(Number int,FullName string,Position string){
	tmp:=d.Tail.Prev
	tmp.Next=&Node{Number: Number,FullName: FullName,Position: Position}
	tmp.Next.Prev=tmp
	tmp.Next.Next=d.Tail
	d.Tail.Prev=tmp.Next
	d.len++
}
func (d *DoubleLinkedList) PopFront(){
	if d.Head.Next==d.Tail{
		fmt.Println("Nothing to delete!")
	}else{
		d.Head.Next=d.Head.Next.Next
		d.Head.Next.Prev=d.Head
		d.len--
	}
}
func (d *DoubleLinkedList) PopBack() {
	if d.Head.Next == d.Tail {
		fmt.Println("Nothing to delete!")
	} else {
		d.Tail.Prev.Prev.Next=d.Tail
		d.Tail.Prev=d.Tail.Prev.Prev
		d.len--
	}
}
func (d *DoubleLinkedList) Begin() Iterator{
	var i Iterator
	i.Node=d.Head.Next
	return i
}
func (d *DoubleLinkedList) End() Iterator {
	i:=Iterator{Node: d.Head}
	for i.HasNext(){
		i.Next()
	}
	i.Prev()
	return i
}
func (d *DoubleLinkedList) Reverse() {
	if d.Head.Next == d.Tail  {
		fmt.Println("Nothing to reverse!")
		return
	}
	if d.Head.Next.Next == d.Tail {
		fmt.Println("Nothing to reverse!")
		return
	}else{
		tmp:=d.Head.Next
		k:=tmp
		for tmp!=d.Tail {
			k = tmp.Next
			tmp.Next = tmp.Prev
			tmp.Prev = k
			tmp=tmp.Prev
		}
		d.Head.Next.Next=d.Tail
		k=d.Head.Next
		d.Head.Next=d.Tail.Prev
		d.Tail.Prev=k
		d.Head.Next.Prev=d.Head
	}
}
func (d *DoubleLinkedList) Sort() {
	if d.Head.Next == d.Tail {
		fmt.Println("Nothing to sort!")
		return
	}
	if d.Head.Next.Next == d.Tail {
		fmt.Println("Nothing to sort!")
		return
	} else {
		tmp := d.Head.Next
		for i:=0;i<d.len;i++{
			for tmp.Next != d.Tail {
				if tmp.Number > tmp.Next.Number {
					b := tmp.Next
					c := tmp
					x := tmp.Next.Next

					if tmp.Prev==nil{
						d.Head.Next=b
						d.Head.Next.Prev=nil
					}else {
						tmp.Prev.Next = b
						tmp = tmp.Prev
						tmp.Next.Prev = tmp
					}
					tmp.Next.Next = c
					tmp = tmp.Next
					tmp.Next.Prev = tmp
					tmp.Next.Next = x
					tmp.Next.Next.Prev = tmp.Next
				}
				tmp = tmp.Next
			}
			tmp = d.Head.Next
		}
	}
}
func (d *DoubleLinkedList) Size() {
	fmt.Println("Size: ", d.len)
}
func (d *DoubleLinkedList) Find(info string) Iterator{
	i:=Iterator{Node: d.Head.Next}
	find:=false
	for i.HasNext(){
		if info==i.Node.FullName{
			find=true
			break
		}
		if info==i.Node.Position{
			find=true
			break
		}
		i.Next()
	}
	if find!=true{
		i.Node=nil
		fmt.Println("There is no such element!")
	}
	return i
}

func main() {
	head:=&Node{Next: nil,Prev: nil}
	tail:=&Node{Prev: head}
	head.Next=tail
	zxc:=DoubleLinkedList{Head:head,Tail: tail}

	zxc.PopFront()
	zxc.PopBack()
	zxc.Print()
	zxc.Reverse()

	zxc.PushBack(8888,"google","it")
	zxc.PushBack(11,"ghul","mider")
	zxc.PushBack(1000,"TROLL","skam")
	zxc.PopFront()
	zxc.Print()

	zxc.PushFront(52,"kunka","ship")
	zxc.PushFront(400,"Andrew","Del")
	zxc.PushFront(5,"Rat","Heist")
	zxc.PushBack(7070,"uwe","lolsz")
	zxc.PushBack(63,"HOUSE","SQUARE")
	zxc.PushFront(55,"kail","carry")
	zxc.PushFront(100,"Mega rat","Hacker")
	zxc.Print()

	zxc.Reverse()
	zxc.Print()

	zxc.Sort()
	zxc.Print()

	u:=zxc.Begin()
	fmt.Println("Begin: ",u.Node)
	u=zxc.End()
	fmt.Println("End: ",u.Node)
	u=zxc.Find("ship")
	fmt.Println("Find: : ",u.Node)
}