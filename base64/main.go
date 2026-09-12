package main

import "fmt"

func main() {
	var w string
	fmt.Print("Give a word: ")
	fmt.Scan(&w)
	var Result string

	// Base64 index table (moved outside the loop)
	base64Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	// STEP 1
	bytes := []byte(w)

	for i, b := range bytes {
		fmt.Printf("\n'%c' = %d = %08b\n", w[i], b, b)
	}

	fmt.Print("\nFull binary: ")
	for i, b := range bytes {
		fmt.Printf("%08b", b)
		if i < len(bytes)-1 {
			fmt.Print(" ")
		}
	}

	// STEP 2
	bytesLen := len(bytes)
	fmt.Printf("\nTotal bytes: %d\n", bytesLen)

	for grp := 0; grp < (bytesLen+2)/3; grp++ {
		start := grp * 3

		if start >= bytesLen {
			break
		}
		b1 := bytes[start]
		var b2, b3 byte = 0, 0
		if start+1 < bytesLen {
			b2 = bytes[start+1]
		}
		if start+2 < bytesLen {
			b3 = bytes[start+2]
		}

		x := (uint32(b1) << 16) | (uint32(b2) << 8) | uint32(b3)
		fmt.Printf("Group %d combined 24-bit: %024b\n", grp+1, x)

		// STEP 3
		fmt.Printf("Group %d 6-bit chunks: ", grp+1)

		chunk1 := (x >> 18) & 0x3F
		chunk2 := (x >> 12) & 0x3F
		chunk3 := (x >> 6) & 0x3F
		chunk4 := x & 0x3F

		fmt.Printf("%06b %06b %06b %06b\n", chunk1, chunk2, chunk3, chunk4)
		fmt.Printf("Decimal values: %d %d %d %d\n", chunk1, chunk2, chunk3, chunk4)

		// STEP 4
		c1 := base64Chars[chunk1]
		c2 := base64Chars[chunk2]
		c3 := base64Chars[chunk3]
		c4 := base64Chars[chunk4]

		Result += string(c1) + string(c2) + string(c3) + string(c4)
		fmt.Printf("Base64 chars: %c %c %c %c\n\n", c1, c2, c3, c4)
	}

	// Padding
	remain := bytesLen % 3
	if remain == 1 {
		Result = Result[:len(Result)-2] + "=="
	} else if remain == 2 {
		Result = Result[:len(Result)-1] + "="
	}

	fmt.Print("\n===[{ FINAL RESULT }]===\n\n")
	fmt.Printf("%s ===> %s\n", w, Result)
}
