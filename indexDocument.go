package main

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"os"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func addProduct(product Product, client *typesense.Client) {

	_, err := client.Collection("products").Documents().Create(product)
	if err != nil {
		fmt.Println(err)
	}
}

func seedProducts(client *typesense.Client) {
	batchSize := 10000
	action := "create"
	var params = &api.ImportDocumentsParams{
		BatchSize: &batchSize,
		Action:    &action,
	}
	importBody, _ := os.Open("seed.jsonl")
	// defer close, error handling ...

	_, err := client.Collection("products").Documents().ImportJsonl(importBody, params)
	if err != nil {
		fmt.Println(err)
	}
}
