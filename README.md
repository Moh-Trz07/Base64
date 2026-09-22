# Base64 From Scratch (Go)

A visual, step-by-step implementation of the **Base64 algorithm** in pure Go — both **encoder** and **decoder**, no `encoding/base64` used under the hood. Every stage (binary → 24-bit groups → 6-bit chunks → Base64 chars → padding) is printed with binary output so you can *see* exactly what's happening.

> ⚠️ **Educational project only.** In real Go code, always use the standard library:
> ```go
> base64.StdEncoding.EncodeToString([]byte("your text"))
> base64.StdEncoding.DecodeString("eW91ciB0ZXh0")
> ```

## Why?

Base64 converts binary data into ASCII text using 64 safe characters (`A–Z`, `a–z`, `0–9`, `+`, `/`). It's used in email attachments, data URLs, JWTs, and embedding images in HTML/CSS. Every 3 bytes (24 bits) become 4 Base64 characters (4 × 6 bits), with `=` padding when the input length isn't a multiple of 3.

This repo exists to **demystify the algorithm** — most people use Base64 as a black box. Here you can watch it work, bit by bit, in both directions.

## How to Run

```bash
go run main.go
```
You'll see a menu:
```
==============================
      BASE64 TOOL (Go)        
==============================
1. Encode
2. Decode
3. Exit
Choose an option:
```

# Examples
## Encoding

```
Give a word: salamoli

's' = 115 = 01110011
'a' = 97  = 01100001
'l' = 108 = 01101100
'a' = 97  = 01100001
'm' = 109 = 01101101
'o' = 111 = 01101111
'l' = 108 = 01101100
'i' = 105 = 01101001

Full binary: 01110011 01100001 01101100 01100001 01101101 01101111 01101100 01101001
Total bytes: 8

Group 1 combined 24-bit: 011100110110000101101100
Group 1 6-bit chunks: 011100 110110 000101 101100
Decimal values: 28 54 5 44
Base64 chars: c 2 F s

Group 2 combined 24-bit: 011000010110110101101111
Group 2 6-bit chunks: 011000 010110 110101 101111
Decimal values: 24 22 53 47
Base64 chars: Y W 1 v

Group 3 combined 24-bit: 011011000110100100000000
Group 3 6-bit chunks: 011011 000110 100100 000000
Decimal values: 27 6 36 0
Base64 chars: b G k A

===[{ FINAL RESULT }]===

salamoli ===> c2FsYW1vbGk=
```

Verified against Go's standard library:

```go
base64.StdEncoding.EncodeToString([]byte("salamoli"))
// → "c2FsYW1vbGk="  ✅
```

## What You'll Learn

- Bit shifting (`<<`, `>>`) and masking (`&`) in Go
- Combining bytes into a 32-bit integer
- Why Base64 inflates data by ~33%
- Why padding exists

## Status

- [x] Encoder (steps 1–5)
- [ ] Decoder
- [ ] Unit tests

## License

MIT
