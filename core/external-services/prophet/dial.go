package prophet

import (
	"buf.build/gen/go/materials-resources/prophet/grpc/go/inventory/v1/inventoryv1grpc"
	"google.golang.org/grpc"
	"log"
)

var client inventoryv1grpc.InventoryServiceClient

func Dial() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("prophet:50058", opts...)
	if err != nil {
		log.Fatalf("there was an issue connecting to the prophet service: %v", err)
	}
	client = inventoryv1grpc.NewInventoryServiceClient(conn)
	return
}
