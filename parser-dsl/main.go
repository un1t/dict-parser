package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(os.Stdout)
	w.Comma = ';'
	defer w.Flush()

	scanner := bufio.NewScanner(file)

	var direction = "ru-sr"
	var word string
	var translation []string
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "\t") {
			line = strings.Replace(line, "[m2", "    [m2", 1)
			var re = regexp.MustCompile(`\[.+?\]`)
			line := re.ReplaceAllString(line, "")

			translation = append(translation, strings.Trim(line, "\t"))
		} else {
			if i > 4 {
				w.Write([]string{direction, word, strings.Join(translation, "\n")})
			}

			word = line
			translation = []string{}
			i += 1
		}
	}
	w.Write([]string{direction, word, strings.Join(translation, "\n")})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
