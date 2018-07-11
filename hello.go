package main

import "fmt"

func main() {
	var num1 int =100
	var num2 int =200
	var ret int
	ret = max(num1,num2)
    fmt.Printf("最大值是 : %d \n",ret)
    a, b := swap("james","cRonalord")
   //  a, b := swap("Mahesh", "Kumar")
    fmt.Println(a,b)
}
func max(num1,num2 int) int {
	if (num1>=num2 ){
		return num1
	}else {
		return num2
	}
}
func swap(x,y string) (string,string) {
	return y,x
}

