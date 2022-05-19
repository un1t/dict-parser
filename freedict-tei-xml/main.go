package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type TEI struct {
	XMLName xml.Name `xml:"TEI"`
	Text    Text     `xml:"text"`
}

type Text struct {
	Body Body `xml:"body"`
}

type Body struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Form    Form    `xml:"form"`
	GramGrp GramGrp `xml:"gramGrp"`
	Sense   []Sense `xml:"sense"`
}

type Form struct {
	Orth string `xml:"orth"`
	Pron string `xml:"pron"`
}

type GramGrp struct {
	Pos string `xml:"pos"`
}

type Sense struct {
	Pos string `xml:"pos"`
}

func main() {
	filename := os.Args[1]
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tei TEI
	xml.Unmarshal(byteValue, &tei)

	for i, entry := range tei.Text.Body.Entries {
		fmt.Println(i, entry.Form.Orth, entry.GramGrp.Pos)
	}
}
