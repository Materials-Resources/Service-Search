package prophet

import (
	inventoryV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/inventory/v1"
	"context"
	"fmt"
	"log"
	"time"
)

func GetProductGroups() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product, err := client.GetProduct(ctx, &inventoryV1.GetProductRequest{InvMastId: 24918})
	if err != nil {
		log.Println(err)
	}
	fmt.Print(product)
}

// GetProductsByGroup Retrieves all products assigned to a specific product group
func GetProductsByGroup(productGroups []string, invMastUid *int32) *inventoryV1.GetProductsByGroupResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.GetProductsByGroup(ctx,
		&inventoryV1.GetProductsByGroupRequest{
			InvMastUid:   *invMastUid,
			ProductGroup: productGroups,
		},
	)
	if err != nil {
		log.Println(err)
	}

	return res
}
