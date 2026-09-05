package main

import "fmt"

func main(){
	var w string
	fmt.Print("Give a word: ")
	fmt.Scan(&w)
    
	// STEP 1 
	fmt.Print("\n===[{ Step 1: Binary Data }]===\n")
    bytes := []byte(w)

	// each character with its decimal and binary
	for i, b := range bytes{
		fmt.Printf("\n'%c' = %d = %08b\n", w[i], b, b)
	}

	fmt.Print("\nFull binary: ")
	for i, b := range bytes {
		fmt.Printf("%08b", b)
		if i < len(bytes)-1 {
			fmt.Print(" ") // adding spaces between bytes just for readability
		}
	}

    // STEP 2
	fmt.Print("\n\n===[{ Step 2: Group into 24-bit Chunks }]===\n\n")
	bytesLen := len(bytes)
	fmt.Printf("Total bytes: %d\n", bytesLen)
	
    // Calculate number of groups needed
    for grp := 0; grp < (bytesLen+2)/3; grp++{
		start := grp * 3

		if start >= bytesLen {break} // check if we have enough bytes for this group
		b1 := bytes[start] 
		var b2, b3 byte = 0, 0
		if start+1 < bytesLen{
			b2 = bytes[start+1]
		}
		if start+2 < bytesLen{
			b3 = bytes[start+2]
		}

		// combine 3 bytes (24 bits) into a single 32-bit integer
		x := (uint32(b1) << 16) | (uint32(b2) << 8) | uint32(b3)

		fmt.Printf("Group %d combined 24-bit: %024b\n", grp+1, x)
	
	// STEP 3
	var sixBitGrps []uint32
    fmt.Printf("Group %d 6-bit chunks: ", grp+1)

	chunk1 := (x >> 18) & 0x3F // First 6 bits (bits 23-18)
	chunk2 := (x >> 12) & 0x3F // Second 6 bits (bits 17-12)
	chunk3 := (x >> 6) & 0x3F  // Third 6 bits (bits 11-6)
	chunk4 := x & 0x3F         // Last 6 bits (bits 5-0)

	sixBitGrps = append(sixBitGrps, chunk1, chunk2, chunk3, chunk4)

		fmt.Printf("%06b %06b %06b %06b\n", chunk1, chunk2, chunk3, chunk4)
		fmt.Printf("Decimal values: %d %d %d %d\n", chunk1, chunk2, chunk3, chunk4)
	}
}
