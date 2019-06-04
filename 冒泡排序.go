package main

import "fmt"

func main(){
	var num []int;
	var long,k int;
	fmt.Println("请输入数组长度：")
	fmt.Scanf("%d",&long)
	fmt.Println("请输入数组数据：")
	for i:=0;i<long;i++ {
		fmt.Scanf("%d",&k)
		num=append(num,k)
	}
    var is=true
	for l:=0;l<long&&is==true;l++ {
		is=false
		for m:=0;m+1<long-l;m++  {
			if(num[m]>num[m+1]){
				is=true
				num[m],num[m+1]=num[m+1],num[m]
			}
		}
	}

	for s:=0;s<long;s++ {
		fmt.Println(num)
	}


}