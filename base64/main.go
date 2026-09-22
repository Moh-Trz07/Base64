/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Give a word: ")
	reader := bufio.NewReader(os.Stdin)
	w, _ := reader.ReadString('\n') // Read until Enter
	w = strings.TrimRight(w, "\r\n") // Remove trailing newline

	var Result string
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
	if bytesLen == 0 {
		fmt.Println("No input given!")
		return
	}
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
	fmt.Printf("%q ===> %s\n", w, Result)  
}*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	fmt.Print("Give a Base64 string : ")
	reader := bufio.NewReader(os.Stdin)
	w, _ := reader.ReadString('\n')
	w = strings.TrimRight(w, "\r\n")

	base64Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	paddingCount := strings.Count(w, "=")

	// STEP 1
	var SixBitVal []uint32

	for _, ch := range w {
		if ch == '='{ continue }
		idx := strings.IndexRune(base64Chars, ch)
		if idx == -1 {
			fmt.Printf("Invaled char: %c\n", ch)
			return
		}
		SixBitVal = append(SixBitVal, uint32(idx))
		fmt.Printf("'%c' -> %d -> %06b\n", ch, idx, idx)
	}
	
	// STEP 2
	var bytesRes []byte
	for i := 0; i < len(SixBitVal); i +=4{
		v1 := SixBitVal[i]
		var v2,v3,v4 uint32 = 0,0,0
		if i+1 < len(SixBitVal){ v2 = SixBitVal[i+1] }
		if i+2 < len(SixBitVal){ v3 = SixBitVal[i+2] }
		if i+3 < len(SixBitVal){ v4 = SixBitVal[i+3] }

		comb := (v1 << 18) | (v2 << 12) | (v3 << 6) | v4
		fmt.Printf("\nGroup %d: %d %d %d %d\n", i/4+1, v1, v2, v3, v4)
		fmt.Printf("Combined 24-bit: %024b\n", comb)

		b1 := byte((comb >> 16) & 0xFF) // 23-16 bits
		b2 := byte((comb >> 8) & 0xFF) // 15-8 bits
		b3 := byte(comb  & 0xFF)  // 7-0 bits

		fmt.Printf(" Bytes: %08b %08b %08b → %d %d %d\n\n", b1, b2, b3, b1, b2, b3)
		bytesRes = append(bytesRes, b1, b2, b3)
	}
    
	// STEP 3
	if paddingCount == 0 { fmt.Println("No padding found") } else {
		
    fmt.Printf("Padding count: %d\n", paddingCount)
	fmt.Printf("Bytes before removal: %v (%d bytes)\n", bytesRes, len(bytesRes))

	if paddingCount > 0 && paddingCount <= len(bytesRes) {
		bytesRes = bytesRes[:len(bytesRes)-paddingCount]
	}
	fmt.Printf("Bytes after removal:  %v (%d bytes)\n\n", bytesRes, len(bytesRes))
}

	
    fmt.Print("===[{ FINAL RESULT }]===\n\n")
	fmt.Printf("%s ===> %s\n", w, string(bytesRes))
}
