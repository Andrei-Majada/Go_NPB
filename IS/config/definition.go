package main

import(
	is "NPB-GO/IS"
	"fmt"
	"os"
)

func verify(args[] string, typeB *int, classp *string){

	*classp = args[2]
	if *classp != "U"{
		if (*classp != "S" && *classp != "W" && *classp != "A" && *classp != "B" && *classp != "C" && *classp != "D" && *classp != "E" && *classp != "F"){
			fmt.Println("definition: Unknown benchmark class ",*classp)
			fmt.Println("definition: Allowed classes are S, W, A, B, C, D, E and F.")
			os.Exit(1)
		}
	}
	if (args[1] == "IS" || args[1] == "is") {
		*typeB = 5
	}else{
		fmt.Println("input error: type 'make help' for more info")
		os.Exit(1)
	}
}

func main(){
	
	var typeB int
	var class string
	
	if len(os.Args) != 3 {
		fmt.Println("input error: type 'make help' for more info")
	}
	args := os.Args
	
	verify(args,&typeB,&class)

	var m int
	if class == "S"{
		m = 24
	}else if class == "W" {
		m = 25
	}else if class == "A" {
		m = 28
	}else if class == "B" {
		m = 30 
	}else if class == "C" {
		m = 32
	}else if class == "D" {
		m = 36
	}else if class == "E" {
		m = 40
	}else {
		fmt.Println("input error: type 'make help' for more info")
		os.Exit(1)
	}

	is.IS(m)
}