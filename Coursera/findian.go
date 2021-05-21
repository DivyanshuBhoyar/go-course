package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter the string")
	fmt.Scan(&input) // aage wala dhokla yaha string enter karega

	//ab apan check karega ki wo "i" se start and "n se end hona chahiye"
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") {
		// Agar yes, toh ye check karege ki usme kahi "a" hai kya
		if strings.IndexAny(input, "a") > 0 { // yeh jo "indexAny" hai , wo agar "a" araha toh positive value return karta, else -1
			fmt.Println("Match found")
		} else { // a nahi mila be
			fmt.Println("Nt found")
		}
	}

}

//samlja kya ?
