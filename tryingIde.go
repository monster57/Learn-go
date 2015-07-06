package main

import(
	"fmt"
	"os" 
	"strings"
)

func main() {
	who:= "World!"
	fmt.Println(len(os.Args))
	if len(os.Args)>1 {
		
		who = strings.Join(os.Args[1:] ," ")
	}

	fmt.Println("Hello" , who)
}