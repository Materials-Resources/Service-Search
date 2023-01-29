package search

import (
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"log"
)

func Products(client *typesense.Client, query string) (results ProductSearchResults) {
	facets := "product_group_id"
	prefix := "false"
	perPage := 20

	params := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: "item_desc",
		FacetBy: &facets,
		Prefix:  &prefix,
		PerPage: &perPage,
	}

	searchRes, err := client.Collection("Products").Documents().Search(params)
	if err != nil {
		log.Println(err)
	}

	for _, hit := range *(searchRes.Hits) {
		results.products = append(results.products, &resultsProduct{
			invMastUid:     (*hit.Document)["id"].(string),
			itemId:         (*hit.Document)["item_id"].(string),
			itemDesc:       (*hit.Document)["item_desc"].(string),
			productGroupId: (*hit.Document)["product_group_id"].(string),
		})
	}
	return results
}
