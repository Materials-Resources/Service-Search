package main

import (
	"github.com/typesense/typesense-go/typesense"
)

func main() {
	client := typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"))

	client.Collection("products").Delete()
	createProductSchema(client)

	seedProducts(client)
}
