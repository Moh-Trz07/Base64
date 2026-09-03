package main

import "fmt"

func main(){
	var w string
	fmt.Print("Give a word: ")
	fmt.Scan(&w)

	fmt.Print("\n{ Step 1: Binary Data }\n")
    bytes := []byte(w)
	for i, b := range bytes{
		fmt.Printf("'%c' = %d = %08b\n", w[i], b, b)
	}

	fmt.Print("\nFull binary: ")
	for i, b := range bytes {
		fmt.Printf("%08b", b)
		if i < len(bytes)-1 {
			fmt.Print(" ") // space between bytes for readability
		}
	}
}
