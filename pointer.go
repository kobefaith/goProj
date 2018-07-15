package main

import "fmt"

func main(){
	var a int =20
	var ip *int 
	ip = &a  
    
	fmt.Printf("a 的地址是: %x \n",&a)

	fmt.Printf("ip 变量存储的指针地址是: %x \n",ip)

	fmt.Printf("*ip 变量的值: %d\n", *ip)	

	var pptr **int
	pptr = &ip
	fmt.Printf("pptr 变量的值: %d \n",**pptr)

	var t int =10
	var b int = 30
	swap(&t, &b)
	fmt.Printf("交换后的值为: %d , %d\n",t,b)	
	t,b = b,t
	fmt.Printf("交换后的值为: %d , %d\n",t,b)	
}
func swap(x *int, y *int){	
	var temp int
	temp = *x
	*x = *y
	*y = temp	
}