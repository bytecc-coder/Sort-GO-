package main

import "fmt"

func shellSort(nums []int)  {
	//计算步长
	for step := len(nums)/3+1; step >0; step=step/3+1{
		//开始插入排序
		for i := step; i < len(nums); i++ {
			//满足条件则插入
			for j:=i;j>=step&&nums[j]<nums[j-step];j=j-step{
				nums[j],nums[j-step]=nums[j-step],nums[j]
			}
		}
		//当步长为1时结束排序
		if step==1{
			break
		}
	}
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
		shellSort(a)
	fmt.Println(a)
}