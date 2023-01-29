package search

import (
	"buf.build/gen/go/materials-resources/search/grpc/go/search/v1/searchv1grpc"
	searchV1 "buf.build/gen/go/materials-resources/search/protocolbuffers/go/search/v1"
	"context"
	"github.com/materials-resources/Service-Search/core/typesense/search"
	"github.com/typesense/typesense-go/typesense"
)

type Server struct {
	searchv1grpc.UnimplementedSearchServiceServer
	TsClient *typesense.Client
}

func (s *Server) SearchTopFive(ctx context.Context, request *searchV1.SearchTopFiveRequest) (*searchV1.SearchTopFiveResponse, error) {

	return nil, nil
}
func (s *Server) SearchAll(ctx context.Context, request *searchV1.SearchAllRequest) (*searchV1.SearchAllResponse, error) {
	products := search.Products(s.TsClient, request.Query)
	results := products.ToGrpc()
	return results, nil
}
