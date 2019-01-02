package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type xmldoc struct {
	XMLName xml.Name `xml:"points"`
	Data []first_test `xml:"first_test"`
}

/* for some reason should capitalise the props here */
type first_test struct {
	Desc string `xml:"desc"`
	XPos string `xml:"xPos"`
	YPos string `xml:"yPos"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) xmldoc {
	dat, err := ioutil.ReadFile(fileName)
	check(err)

	parsedData := xmldoc{}

	fmt.Printf(string(dat))

	xml.Unmarshal(dat, &parsedData)
	fmt.Printf("XMLName: %#v\n", parsedData.XMLName)
	//fmt.Printf("desc: %q\n", parsedData.desc)
	return parsedData
}
