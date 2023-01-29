package models

type Product struct {
	ID             string `json:"id"`
	ItemId         string `json:"item_id"`
	ItemDesc       string `json:"item_desc"`
	ProductGroupId string `json:"product_group_id"`
}
