package main

import "fmt"

//斐波那契数列

func main()  {
	var i int  =1;
	var j int  =1;
	var m int;
	var arr[] int;
	arr=append(arr,1,1)
    fmt.Println("输入斐波那契数上限：")
	fmt.Scanf("%d",&m)
	for k:=3;;k++{
		if i+j>m{
			break
		}
		arr=append(arr,i+j)
		i,j=j,i
		j=j+i

	}
	fmt.Println(arr)
}
