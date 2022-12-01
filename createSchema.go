package main

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func createProductSchema(client *typesense.Client) {
	schema := &api.CollectionSchema{
		Name: "products",
		Fields: []api.Field{
			{
				Name: "inv_mast_uid",
				Type: "int32",
			},
			{
				Name: "item_desc",
				Type: "string",
			},
			{
				Name: "item_id",
				Type: "string",
			},
			{
				Name: "product_group_id",
				Type: "string",
			},
		},
	}
	_, err := client.Collections().Create(schema)
	if err != nil {
		fmt.Println(err)
		return
	}
}
