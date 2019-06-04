package main

import (
	"fmt"
	"math"
)

func main()  {
	var c float64=0
	var i int=1
	var j int =0
	for ;c<88480;i++{
		c=math.Pow(2,float64(i))
		j++
	}

	fmt.Println(j)
}
