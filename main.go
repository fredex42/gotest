package main

import "flag"
import "fmt"

func main () {
	fmt.Println("Hello world")
	fileName := flag.String("file","test1.xml","XML file to read")

	flag.Parse()

	testData := readFile(*fileName)
	fmt.Println(testData)
}
