package main

import "fmt"
func partition(array []int, left, right int) int {
	//baseNum 保存轴值
	baseNum := array[left]
	//当数组左又下标未相遇时
	for left < right{
		//从右往左遍历找到一个比轴值小的值
		for (array[right] >= baseNum && right > left){
			right--
		}
		array[left] = array[right]
		//从左往右遍历找到一个比轴值大的值
		for (array[left] <=baseNum && right > left) {
			left++
		}
		array[right] = array[left]
	}
	//把轴值回填到分界位置上
	array[right] = baseNum
	return right
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	//分割后轴值的正确位置
	index := partition(array,left,right)
	//递归排序左半部分
	quickSort(array, left, index - 1)
	//递归排序右半部分
	quickSort(array, index + 1, right)
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
	quickSort(a,0, len(a)-1)
	fmt.Println(a)
}
