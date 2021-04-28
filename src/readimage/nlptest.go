package main

import (
	"context"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func nlpTest(text string) {
	ctx := context.Background()

	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	//text := "Hello, world! You are so beautiful. You gave me America, home of the free."

	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text for sentiment analysis: %v", err)
	}

	//fmt.Printf("Text: %v\n", text)

	fmt.Println("*******Sentiment:Score:", sentiment.DocumentSentiment.Score, ", Magnitude:", sentiment.DocumentSentiment.Magnitude)

	entity, err := client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text for entity analysis: %v", err)
	}
	fmt.Println("******Entities: ", entity.String())

	syntax, err := client.AnnotateText(ctx, &languagepb.AnnotateTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		Features: &languagepb.AnnotateTextRequest_Features{
			ExtractSyntax: true,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text for annotation analysis: %v", err)
	}

	//fmt.Println("******Syntax - Tokens: ", syntax.Tokens)
	for _, token := range syntax.Tokens {
		fmt.Println("***TOKEN: LABEL-", token.DependencyEdge.Label.String(), ", TEXT-", token.Text.Content, ", PART OF SPEECH-", token.PartOfSpeech.String())
	}
	//fmt.Println("******Syntax - Sentences: ", syntax.Sentences)
	for _, sentence := range syntax.Sentences {
		fmt.Println("***SENTENCE", sentence.Text.GetContent())
	}
	//fmt.Println("******Syntax - String: ", syntax.String())

}
