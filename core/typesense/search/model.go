package search

import searchV1 "buf.build/gen/go/materials-resources/search/protocolbuffers/go/search/v1"

type ProductSearchResults struct {
	facets   resultsFacets
	products []*resultsProduct
}

type resultsFacets struct {
	productGroups []resultsProductGroupFacet
}

type resultsProductGroupFacet struct {
	name   string
	count  int32
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
