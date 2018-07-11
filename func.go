package main
import "fmt"

func getSequence() func() int {
	i :=0
	return func() int {
	  i+=1
	  return i
	}
}
func main(){
	nextNum := getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	nextNum1 :=getSequence()
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())

}