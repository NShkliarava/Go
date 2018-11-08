package main 

import"fmt"

func main() {

	var a[5] float32
		a[0]=8
		a[1]=7
		a[2]=9
		a[3]=6
		a[4]=5
	var total float32 = 0
		for i:=0; i<5; i++ {
			total += a[i]
	}
	fmt.Println("average",total/5)
	fmt.Println("max", a[2])
	fmt.Println("min", a[4])
	}