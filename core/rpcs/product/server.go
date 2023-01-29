package product

import (
	"buf.build/gen/go/materials-resources/search/grpc/go/product/v1/productv1grpc"
	productV1 "buf.build/gen/go/materials-resources/search/protocolbuffers/go/product/v1"
	"context"
	"github.com/materials-resources/Service-Search/core/typesense/index"
	"github.com/typesense/typesense-go/typesense"
)

type Server struct {
	productv1grpc.UnimplementedProductServiceServer
	TsClient *typesense.Client
}

func (s *Server) IndexProducts(ctx context.Context, request *productV1.IndexProductsRequest) (response *productV1.IndexProductsResponse, err error) {
	err = index.Product(s.TsClient)
	return nil, nil
}
