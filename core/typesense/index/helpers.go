package index

import (
	inventoryV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/inventory/v1"
	"strconv"
)

func convertProphetProductsToDoc(products []*inventoryV1.GetProductsByGroupResponseProduct) []interface{} {

	var docs []interface{}

	for _, product := range products {
		docs = append(docs, &struct {
			InvMastUid     string `json:"id"`
			ItemId         string `json:"item_id"`
			ItemDesc       string `json:"item_desc"`
			ProductGroupId string `json:"product_group_id"`
		}{
			InvMastUid:     strconv.Itoa(int(product.InvMastUid)),
			ItemId:         product.ItemId,
			ItemDesc:       product.ItemDesc,
			ProductGroupId: product.ProductGroupId,
		})
	}

	return docs
}
