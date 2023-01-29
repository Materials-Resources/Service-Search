package schemas

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func registerProductGroupSchema(client *typesense.Client) error {
	schema := &api.CollectionSchema{
		Name: "Product Groups",
		Fields: []api.Field{
			{
				Name: "product_group_id",
				Type: "string",
			},
			{
				Name: "product_group_desc",
				Type: "string",
			},
		}}
	_, err := client.Collections().Create(schema)
	if err != nil {
		return fmt.Errorf("failed to create Product Groups schema: %w", err)
	}
	return nil
}
