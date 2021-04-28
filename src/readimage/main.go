package main

import (
	"fmt"
	"log"
)

func main() {
	filename := "C:\\govision\\receipt.jpg"
	//filename := "C:\\govision\\iFly.PNG"
	ocrText, err := visionTest(filename)
	if err != nil {
		log.Fatalf("visionTest Failed: %v", err)
	}
	fmt.Println(ocrText)

	nlpTest(ocrText)

}
