package main

import (
   "fmt"
   "os"
   "bufio"  
)
func main(){	
	if len(os.Args) >=2 {
      for _, arg := range os.Args[1:]{
		  fmt.Printf("%s\n", arg)
      }
	}else {			
		input := bufio.NewScanner(os.Stdin)
	    for input.Scan() {
		   fmt.Printf("%s\n", input.Text())
	    }
	}	
}
