package main

import (
	"fmt"
)

func sort(a[] int) (k int,l int){
	s,t,p,q:=0,0,0,a[0]
	for i:=0;i<len(a);i++{
		if(a[i]>=t){
			t=a[i]
			s=i
		}
		if(a[i]<q){
			q=a[i]
			p=i
		}
	}
	return s, p
}

func main()  {
	var i,x int
	_, _ = fmt.Scanf("%d", &i)
	var c []int
	for j:=0;j<i;j++  {
		_, _ = fmt.Scanf("%d",&x)
		c=append(c,x)
	}
	fmt.Println(c)

}

