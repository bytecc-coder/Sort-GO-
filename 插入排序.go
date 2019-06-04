package main

import "fmt"

func main()  {
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
	for i:=1;i< len(a);i++  {
		for j:=i;j>=1&&a[j]<a[j-1];j=j-1{
			a[j],a[j-1]=a[j-1],a[j]
		}
	}
	fmt.Println(a)

}
