package product

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	GetProduct endpoint.Endpoint
}

func MakeEndpoints(s Service) EndPoints {
	return EndPoints{GetProduct: makeGetProductEndPoints(s)}
}

func makeGetProductEndPoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		product, err := s.GetProduct(ctx, req.id)

		return GetProductResponse{
			Product: product,
		}, err
	}
}
