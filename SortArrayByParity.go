package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	left:=0
	i:=0
	right:=len(nums)-1
	for i<right{
		if nums[i]%2==0{
			nums[left],nums[i]=nums[i],nums[left]
			left++
			i++
		}
		if nums[i]%2==1{
			nums[right],nums[i]=nums[i],nums[right]
			right--
		}
	}
	return nums
}
func main() {
	arr := []int{17, 18, 11, 11, 0, 40, 9481, 2121, 7, 8}
	fmt.Println(arr)
	sortArrayByParity(arr)
	fmt.Println(arr)
}