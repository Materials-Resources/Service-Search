package search

import (
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"log"
	"math"
	"strconv"
)

func Products(client *typesense.Client, query string) (results ProductSearchResults) {
	facets := "product_group_id"
	prefix := "false"
	perPage := 20
	page := 1

	params := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: "item_desc",
		FacetBy: &facets,
		Prefix:  &prefix,
		PerPage: &perPage,
		Page:    &page,
	}

	searchRes, err := client.Collection("Products").Documents().Search(params)
	if err != nil {
		log.Println(err)
	}

	results.currentPage = int32(*(searchRes.Page))

	results.totalPages = int32(math.Ceil(float64(int32(*(searchRes.Found) / searchRes.RequestParams.PerPage))))

	for _, facet := range *(searchRes.FacetCounts) {
		if *facet.FieldName == "product_group_id" {
			for _, productGroupFacet := range *(facet.Counts) {
				results.facets.productGroups = append(results.facets.productGroups, &resultsProductGroupFacet{
					name:  *productGroupFacet.Value,
					count: strconv.Itoa(*productGroupFacet.Count),
				})
			}
		}
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
