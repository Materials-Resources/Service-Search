package product_group

import (
	"buf.build/gen/go/materials-resources/search/grpc/go/productGroup/v1/product_groupv1grpc"
	"github.com/typesense/typesense-go/typesense"
)

type Server struct {
	product_groupv1grpc.ProductGroupServiceServer
	TsClient *typesense.Client
}
