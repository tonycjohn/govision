package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func visionTest(fileName string) (string, error) {

	ctx := context.Background()

	//create client
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	defer client.Close()

	//filename := "C:\\govision\\wakeupcat.jpg"
	//filename := "C:\\govision\\iFly.PNG"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	fmt.Println("*****LABELS*****")
	for _, label := range labels {
		fmt.Println(label.Description)
	}

	fmt.Println("*****TEXT*****")
	texts, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		return "", err
	}
	return texts.GetText(), nil
}
