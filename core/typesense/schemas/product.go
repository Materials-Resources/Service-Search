package schemas

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func registerProductSchema(client *typesense.Client) error {
	facet := true
	schema := &api.CollectionSchema{
		Name: "Products",
		Fields: []api.Field{
			{
				Name: "item_id",
				Type: "string",
			},
			{
				Name: "item_desc",
				Type: "string",
			},
			{
				Name:  "product_group_id",
				Type:  "string",
				Facet: &facet,
			},
		},
	}
	_, err := client.Collections().Create(schema)
	if err != nil {
		return fmt.Errorf("failed to create Products schema: %w", err)
	}
	return nil
}
