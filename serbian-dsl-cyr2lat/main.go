package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const CYR_ALPHABET = "А Б Ц Ћ Ч Д Ђ Џ  Е Ф Г Х И Ј К Л Љ  М Н Њ  О П Р С Ш Т У В З Ж"
const LAT_ALPHABET = "A B C Ć Č D Đ Dž E F G H I J K L Lj M N Nj O P R S Š T U V Z Ž"

func cyr2lat(s string) string {
	cyrAlphabet := strings.Fields(CYR_ALPHABET + " " + strings.ToLower(CYR_ALPHABET))
	latinAlphabet := strings.Fields(LAT_ALPHABET + " " + strings.ToLower(LAT_ALPHABET))

	for i, charFrom := range cyrAlphabet {
		charTo := latinAlphabet[i]
		s = strings.ReplaceAll(s, charFrom, charTo)
	}

	return s
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "\t") {
			line = cyr2lat(line)
		}
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
