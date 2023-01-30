package search

import (
	searchV1 "buf.build/gen/go/materials-resources/search/protocolbuffers/go/search/v1"
)

type ProductSearchResults struct {
	facets      resultsFacets
	products    []*resultsProduct
	totalPages  int32
	currentPage int32
}

type resultsFacets struct {
	productGroups []*resultsProductGroupFacet
}

type resultsProductGroupFacet struct {
	name   string
	count  string
	active bool
}

type resultsProduct struct {
	invMastUid     string
	itemId         string
	itemDesc       string
	productGroupId string
}

func (m *ProductSearchResults) ToGrpc() *searchV1.SearchAllResponse {
	response := new(searchV1.SearchAllResponse)

	response.TotalPages = m.totalPages
	response.CurrentPage = m.currentPage

	response.Facets = &searchV1.SearchAllResponse_Facets{}
	// Append Product Group facets
	for _, productGroup := range m.facets.productGroups {

		response.Facets.ProductGroups = append(response.Facets.ProductGroups, &searchV1.SearchAllResponse_Facets_ProductGroup{
			Name:  productGroup.name,
			Count: productGroup.count,
		})

	}

	// Append product results
	for _, product := range m.products {
		response.SearchResults = append(response.SearchResults, &searchV1.SearchAllResponse_Results{
			InvMastUid:     product.invMastUid,
			ItemId:         product.itemId,
			ItemDesc:       product.itemDesc,
			ProductGroupId: product.productGroupId,
		})
	}
	return response
}
