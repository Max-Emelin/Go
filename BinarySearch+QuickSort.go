package main

import "fmt"

func BinarySearch(array []int, data int ) int{
	left:=0
	midle:=0
	right:=len(array)-1
	for left<=right {
		midle = (left + right) / 2
		if (data < array[midle]) {
			right = midle - 1
		}
		if (data > array[midle]) {
			left = midle + 1
		}
		if (data == array[midle]) {
			return midle
		}
	}
	fmt.Println("No number!")
	return -1
}
func QuickSort(array []int,left int,right int) {
	if len(array) == 0 {
		return
	}
	if left >= right {
		return
	}
	midle := left + (right-left)/2
	opora := array[midle]
	i:=left
	j:=right
	for i<=j{
		for array[i]<opora{
			i++
		}
		for array[j]>opora{
			j--
		}
		if i<=j{
			array[i],array[j]=array[j],array[i]
			i++
			j--
		}
	}
	if left<j{
		QuickSort(array,left,j)
	}
	if right>i{
		QuickSort(array,i,right)
	}
}

func main() {
	arr:=[]int{20,18,16,14,12,10,8,6,4,2}
	fmt.Println(arr)
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
	fmt.Println(BinarySearch(arr,10))
}