// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {	
	input := bufio.NewScanner(os.Stdin)	
    for input.Scan(){
	   fmt.Printf("the input is :  ")		
	   fmt.Println(input.Text())
    }	
}

//!-