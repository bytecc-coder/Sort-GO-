package main

import "fmt"

func MergeSort(nums []int)  {

}

func main() {
	var a []int
	var b int
	var c int
	fmt.Println("请输入待排序列长度")
	_, _ = fmt.Scanf("%d", &b);
	fmt.Println("请输入待排序列：")
	for i:=0;i<b;i++{
		_, _ = fmt.Scanf("%d", &c)
		a=append(a,c)
	}
	fmt.Println(len(a))
	MergeSort(a)
	fmt.Println(a)
}
