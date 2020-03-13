package main

import (
	"fmt"
)

//第一个数字和第二个数字比较，如果大于后边的数字就交换两个数字

/*

func bubble(arr []int){
	for i :=0;i<len(arr)-1;i++{
		for j :=0;j<len(arr)-1-i;j++{
			if arr[j] >arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	
}

*/

func bubble(arr []int){
	for i :=0;i<len(arr)-1;i++{
		for j :=0;j<len(arr)-1-i;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}




func main(){
	ss := []int{2,3,6,1,1,4}
	fmt.Println(ss)
	//bubbleSort(ss)
	bubble(ss)
	fmt.Println(ss)
}
