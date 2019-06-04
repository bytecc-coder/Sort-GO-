package main

import "fmt"

func main()  {
     var num = [5]int{1,3,9,8,5}
     for i:=0;i< len(num);i++{
		 for j:=i+1;j< len(num);j++ {
			 if num[i]>num[j]{
				 num[i],num[j]=num[j],num[i]
			 }
		 }
	 }
     fmt.Println(num)
}
