package main

import (
	"buf.build/gen/go/materials-resources/search/grpc/go/product/v1/productv1grpc"
	"buf.build/gen/go/materials-resources/search/grpc/go/productGroup/v1/product_groupv1grpc"
	"buf.build/gen/go/materials-resources/search/grpc/go/search/v1/searchv1grpc"
	"fmt"
	"github.com/materials-resources/Service-Search/core/external-services/prophet"
	"github.com/materials-resources/Service-Search/core/rpcs/product"
	"github.com/materials-resources/Service-Search/core/rpcs/product-group"
	"github.com/materials-resources/Service-Search/core/rpcs/search"
	"github.com/materials-resources/Service-Search/core/typesense/schemas"
	"github.com/typesense/typesense-go/typesense"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var client *typesense.Client

func main() {

	client = typesense.NewClient(
		typesense.WithServer(os.Getenv("TS_URL")),
		typesense.WithAPIKey(os.Getenv("TS_API_KEY")),
	)
	schemas.RegisterSchemas(client)

	connectToServices()

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
func connectToServices() {
	prophet.Dial()
}

func run() error {
	listenOn := "0.0.0.0:50057"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to start service on: %s %w", listenOn, err)

	}
	server := grpc.NewServer()

	searchv1grpc.RegisterSearchServiceServer(server, &search.Server{TsClient: client})
	productv1grpc.RegisterProductServiceServer(server, &product.Server{TsClient: client})
	product_groupv1grpc.RegisterProductGroupServiceServer(server, &product_group.Server{TsClient: client})

	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to server gRPC server: %w", err)
	}
	return nil
}
