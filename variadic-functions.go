package main

import "fmt"

func sum(nums ...int){
	total :=0
	for _,num :=range nums{
		total +=num
	}
	fmt.Println(total)
}

func main()  {
	sum(1,3)
	sum(1,3,4,67,8,89)

	sums :=[]int{1,3,4,35,34}
	sum(sums...)
}