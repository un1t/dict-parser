package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func DecodeStr(s string) int64 {
	var total int64 = 0
	for i, char := range Reverse(s) {
		total += DecodeChar(char) * int64(math.Pow(64, float64(i)))
	}
	return total
}

func DecodeChar(char rune) int64 {
	if char >= 'A' && char <= 'Z' {
		return int64(char) - 65
	} else if char >= 'a' && char <= 'z' {
		return int64(char) - 71
	} else if char >= '0' && char <= '9' {
		return int64(char) + 4
	} else if char == '+' {
		return 62
	} else if char == '/' {
		return 63
	}

	log.Fatal(fmt.Sprintf("Invalid char: '%s'", string(char)))
	return -1
}

func ReadBytes(f *os.File, offset int64, length int64) []byte {
	f.Seek(offset, 0)
	bs := make([]byte, length)

	_, err := f.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	return bs
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	indexFilename := os.Args[1]
	dictFilename := os.Args[2]

	indexFile, err := os.Open(indexFilename)
	if err != nil {
		panic(err)
	}
	defer indexFile.Close()

	dictFile, err := os.Open(dictFilename)
	if err != nil {
		panic(err)
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(indexFile)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "\t")
		headword := items[0]
		offset := items[1]
		length := items[2]

		fmt.Println("---")
		fmt.Println(headword)
		fmt.Println(string(ReadBytes(dictFile, DecodeStr(offset), int64(DecodeStr(length)))))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bs := ReadBytes(dictFile, 618037, 119)

	fmt.Printf(string(bs))
}
