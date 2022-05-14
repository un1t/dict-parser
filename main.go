package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func DecodeStr(s string) int {
	total := 0
	for i, char := range Reverse(s) {
		total += DecodeChar(char) * int(math.Pow(64, float64(i)))
	}
	return total
}

func DecodeChar(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char) - 65
	} else if char >= 'a' && char <= 'z' {
		return int(char) - 71
	} else if char >= '0' && char <= '9' {
		return int(char) + 4
	} else if char == '+' {
		return 62
	} else if char == '/' {
		return 63
	}

	log.Fatal(fmt.Sprintf("Invalid char: '%s'", string(char)))
	return -1
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(618037, 0)
	bs := make([]byte, 119)

	_, err = f.Read(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(bs))
}
